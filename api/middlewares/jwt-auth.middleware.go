package middleware

import (
	"mit-api/internal/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) == 2 {
			token := t[1]
			isAuthorized, err := helpers.IsAuthorized(token, "my_secret_key")
			if isAuthorized {
				userId, err := helpers.ExtractIDFromToken(token, "my_secret_key")
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
					c.Abort()
					return
				}
				uintVal := uint(userId)
				c.Set("x-user-id", uintVal)
				c.Next()

				return
			}

			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()

			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		c.Abort()
	}
}
