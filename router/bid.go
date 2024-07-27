package router

import (
	"github.com/Just-A-NoobieDev/auction-go-server/internal/bidding"
	"github.com/Just-A-NoobieDev/auction-go-server/middleware"
	"github.com/gin-gonic/gin"
)

func BidRouters(routeGroup *gin.RouterGroup, bidHandler *bidding.BidHandler) {
	bid := routeGroup.Group("/bids")
	bid.Use(middleware.AuthMiddleware())
	{
		bid.POST("/:id", bidHandler.PlaceBid) // manual placement of bid
		bid.GET("/:id", bidHandler.GetBidsByID)
	}
}