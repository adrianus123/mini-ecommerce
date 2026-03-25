package delivery

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

var services = map[string]string{
	"orders":   "http://localhost:8081",
	"payments": "http://localhost:8082",
	"products": "http://localhost:8084",
}

func DynamicProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		service := c.Param("service")

		target, ok := services[service]
		if !ok {
			c.JSON(404, gin.H{"message": "Service not found"})
		}

		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)

		c.Request.Host = url.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
