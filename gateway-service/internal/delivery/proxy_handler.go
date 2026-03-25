package delivery

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyHandler(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)

		c.Request.Host = url.Host
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
