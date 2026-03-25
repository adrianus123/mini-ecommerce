package main

import (
	"os"
	"payment-service/internal/delivery/http/handler"
	"payment-service/internal/infrastructure/database"
	"payment-service/internal/repository"
	"payment-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error load .env file: " + err.Error())
	}

	db, err := database.NewPostgres(os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic("DB connection failed: " + err.Error())
	}

	repo := repository.NewPaymentRepo(db)
	usecase := usecase.NewPaymentUsecase(repo)
	handler := handler.NewPaymentHandler(usecase)

	r := gin.Default()

	api := r.Group("/api/payments")
	api.GET("/test", handler.Test)
	api.POST("/create", handler.CreatePayment)

	r.Run(":8082")
}
