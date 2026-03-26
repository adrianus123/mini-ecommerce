package main

import (
	"context"
	"inventory-service/internal/delivery/kafka"
	"inventory-service/internal/infrastructure/database"
	"inventory-service/internal/repository"
	"inventory-service/internal/usecase"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load .env file: " + err.Error())
	}

	db, err := database.NewPostgres(os.Getenv("DB_CONNECTION"))
	// if err != nil {
	// 	panic("DB connection failed: " + err.Error())
	// }

	ctx := context.Background()
	eventRepo := repository.NewEventRepository()
	productRepo := repository.NewProductRepository(db)
	usecase := usecase.NewEventUseCase(eventRepo, productRepo)

	consumer := kafka.NewConsumer(
		[]string{"localhost:9092"},
		"SYNC_INVENTORY",
		"inventory_service_group",
		usecase,
	)

	consumer.Start(ctx)

	log.Println("Service running...")
}
