package main

import (
	"auth-service/internal/delivery/http/handler"
	"auth-service/internal/infrastructure/auth0"
	"auth-service/internal/infrastructure/database"
	"auth-service/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env " + err.Error())
	}

	db, err := database.NewPostgres(os.Getenv("DB_CONNECTION"))
	if err != nil {
		panic("DB connection failed: " + err.Error())
	}

	repo := database.NewUserRepo(db)

	auth0Client := &auth0.Auth0Client{
		Domain:       os.Getenv("AUTH0_DOMAIN"),
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		Audience:     os.Getenv("AUTH0_AUDIENCE"),
		MgmtToken:    os.Getenv("AUTH0_MGMT_TOKEN"),
	}

	usecase := usecase.NewAuthUsecase(auth0Client, repo)
	handler := handler.NewAuthHandler(usecase)

	r := gin.Default()

	api := r.Group("/api/auth")
	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)

	r.Run(":8083")
}
