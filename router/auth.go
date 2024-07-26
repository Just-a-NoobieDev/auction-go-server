package router

import (
	"github.com/Just-A-NoobieDev/auction-go-server/internal/user"
	"github.com/gin-gonic/gin"
)

func AuthRouters(router *gin.RouterGroup, userHandler *user.UserHandler) {
	auth := router.Group("/auth")
	{
		auth.GET("/", func(c *gin.Context) {
			c.JSON(501, gin.H{
				"message": "Not Implemented",
			})
		})
		auth.POST("/register", userHandler.RegisterUser)
		auth.POST("/login", userHandler.AuthenticateUser)
	}
}