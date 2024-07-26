package router

import (
	"github.com/Just-A-NoobieDev/auction-go-server/internal/user"
	"github.com/Just-A-NoobieDev/auction-go-server/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouters(router *gin.RouterGroup, userHandler *user.UserHandler) {
	user := router.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", userHandler.GetUser)
	}
}