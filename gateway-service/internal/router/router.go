package router

import (
	"gateway-service/internal/delivery"
	"gateway-service/internal/middleware"
	"gateway-service/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, usecase *usecase.RateLimiterUsecase) {
	api := r.Group("/api")
	api.POST("/auth/*path", delivery.ProxyHandler(os.Getenv("AUTH_SERVICE")))

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(os.Getenv("AUTH0_AUDIENCE"), os.Getenv("AUTH0_ISSUER")))
	protected.Use(middleware.RateLimiterMiddleware(usecase, 100, 3600))
	{
		protected.Any("/:service/*path", delivery.DynamicProxy())
	}
}
