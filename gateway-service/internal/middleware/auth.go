package middleware

import (
	"fmt"
	"gateway-service/utils"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwks keyfunc.Keyfunc

func InitJwks(domain string) error {
	jwksUrl := fmt.Sprintf("https://%s/.well-known/jwks.json", domain)

	k, err := keyfunc.NewDefault([]string{jwksUrl})
	if err != nil {
		return err
	}

	jwks = k
	return nil
}

func AuthMiddleware(audience, issuer string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, jwks.Keyfunc)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid token"})
			return
		}

		if !utils.ContainsString(claims.Audience, audience) {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid audience"})
		}

		if claims.Issuer != issuer {
			c.AbortWithStatusJSON(401, gin.H{"message": "Invalid issuer"})
			return
		}

		c.Request.Header.Set("X-User-Id", claims.Subject)

		c.Set("user_id", claims.Subject)
		c.Next()
	}
}
