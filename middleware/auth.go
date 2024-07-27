package middleware

import (
	"errors"
	"strings"

	"github.com/Just-A-NoobieDev/auction-go-server/pkg/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := extractAndValidateToken(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}


		c.Set("userID", claims.UserID.String())
		
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := extractAndValidateToken(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims.UserRole != "admin" {
			c.JSON(403, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID.String())
		c.Next()
	}
}

func extractAndValidateToken (c *gin.Context) (*auth.Claims, error) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return nil, errors.New("Unauthorized")
	}

	tokenString := strings.Split(authHeader, "Bearer ")[1]
	claims, err := auth.ValidateToken(tokenString)

	if err != nil {
		return nil, err
	}

	return claims, nil
}