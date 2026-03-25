package main

import (
	"context"
	"os"
	"product-service/internal/delivery/http/handler"
	"product-service/internal/infrastructure/database"
	"product-service/internal/infrastructure/redis"
	"product-service/internal/repository"
	"product-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load .env file: " + err.Error())
	}

	ctx := context.Background()

	redisClient := redis.NewRedisClient(ctx)
	err = redisClient.Ping(ctx).Err()
	if err != nil {
		panic("Redis connection failed: " + err.Error())
	}

	redisRepository := redis.NewRedisRepository(redisClient, ctx)

	db, err := database.NewPostgres(os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	productRepository := repository.NewProductRepo(db)
	productUsecase := usecase.NewProductUsecase(productRepository, *redisRepository)
	productHandler := handler.NewProductHandler(productUsecase)

	r := gin.Default()
	api := r.Group("/api/products")

	api.GET("/list", productHandler.GetProducts)
	api.GET("/:id", productHandler.GetByID)

	r.Run(":8084")
}
