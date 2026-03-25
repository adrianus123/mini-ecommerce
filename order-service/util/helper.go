package util

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) string {
	userID := c.GetHeader("X-User-Id")
	return userID
}
