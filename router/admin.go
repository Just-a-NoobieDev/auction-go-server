package router

import (
	"github.com/Just-A-NoobieDev/auction-go-server/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouters(router *gin.RouterGroup, handlers *Handlers) {
	admin := router.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	{
		admin.GET("/", handlers.UserHandler.GetAllUsers)
		admin.DELETE("/:id", handlers.UserHandler.DeleteUser)
	}

}