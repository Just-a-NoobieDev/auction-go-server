package router

import (
	"github.com/Just-A-NoobieDev/auction-go-server/internal/auction"
	"github.com/Just-A-NoobieDev/auction-go-server/middleware"
	"github.com/gin-gonic/gin"
)

func AuctionRouters(router *gin.RouterGroup, auctionHandler *auction.AuctionHandler) {
	auction := router.Group("/auction")
	auction.Use(middleware.AuthMiddleware())
	{
		auction.GET("/", auctionHandler.ListAuctions)
		auction.GET("/:id", auctionHandler.GetAuctionByID)
		auction.POST("/", auctionHandler.CreateAuction)
		auction.PUT("/:id", auctionHandler.UpdateAuction)
		auction.DELETE("/:id", auctionHandler.DeleteAuction)
	}
}