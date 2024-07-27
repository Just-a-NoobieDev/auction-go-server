package router

import (
	"github.com/Just-A-NoobieDev/auction-go-server/internal/auction"
	"github.com/Just-A-NoobieDev/auction-go-server/internal/bidding"
	"github.com/Just-A-NoobieDev/auction-go-server/internal/user"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		Engine: gin.Default(),
	}
}

type Handlers struct {
	UserHandler *user.UserHandler
	AuctionHandler *auction.AuctionHandler
	BidHandler *bidding.BidHandler
}

func (r *Router) SetupRouter(handlers *Handlers) {

	r.Engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	mainRouter := r.Engine.Group("/api/v1")

	AuthRouters(mainRouter, handlers.UserHandler)
	UserRouters(mainRouter, handlers.UserHandler)
	AdminRouters(mainRouter, handlers)
	AuctionRouters(mainRouter, handlers.AuctionHandler)
	BidRouters(mainRouter, handlers.BidHandler)
}

