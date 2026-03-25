package main

import (
	"context"
	"gateway-service/internal/infrastructure"
	"gateway-service/internal/middleware"
	"gateway-service/internal/router"
	"gateway-service/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error load .env file: " + err.Error())
	}

	err = middleware.InitJwks(os.Getenv("AUTH0_DOMAIN"))
	if err != nil {
		panic("Failed to init JWKS: " + err.Error())
	}

	redisClient := infrastructure.NewRedisClient()

	err = redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic("Redis connection failed: " + err.Error())
	}

	repo := infrastructure.NewRedisRateLimiterRepository(redisClient)
	usecase := usecase.NewRateLimiterUsecase(repo)

	r := gin.Default()

	r.Use(middleware.Logging())

	router.SetupRouter(r, usecase)

	r.Run(":8080")
}
