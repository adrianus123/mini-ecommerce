package middleware

import (
	"gateway-service/internal/domain"

	"github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(usecase domain.RateLimiterUsecase, limit int, window int) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		var key string

		if exists {
			key = "rate_limit:user:" + userID.(string) + ":" + c.Request.Method + ":" + c.Request.URL.Path
		} else {
			key = "rate_limit:ip:" + c.ClientIP()
		}

		allowed, err := usecase.IsAllowed(key, limit, window)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal error"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(429, gin.H{"message": "Too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
