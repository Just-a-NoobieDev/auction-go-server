package middleware

import (
	"strings"

	"github.com/Just-A-NoobieDev/auction-go-server/pkg/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims, err := auth.ValidateToken(tokenString)

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}