package main

import (
	"order-service/internal/delivery/http/handler"
	"order-service/internal/infrastructure/database"
	"order-service/internal/infrastructure/http"
	"order-service/internal/infrastructure/kafka"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		panic("Error load .env file: " + err.Error())
	}

	db, err := database.NewPostgres(os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic("DB connection failed: " + err.Error())
	}

	producer := kafka.NewKafkaProducer(
		[]string{"localhost:9092"}, "SYNC_INVENTORY",
	)

	kafkaUsecase := usecase.NewPublishEventUsecase(producer)

	productService := http.NewProductClient(os.Getenv("PRODUCT_SERVICE_URL"))
	paymentService := http.NewPaymentClient(os.Getenv("PAYMENT_SERVICE_URL"))

	repo := repository.NewOrderRepo(db)
	usecase := usecase.NewOrderUsecase(repo, productService, paymentService, kafkaUsecase)
	handler := handler.NewOrderHandler(usecase)

	r := gin.Default()

	api := r.Group("/api/orders")
	api.GET("/test", handler.Test)
	api.POST("/save", handler.CreateOrder)

	api.Use()

	r.Run(":8081")
}
