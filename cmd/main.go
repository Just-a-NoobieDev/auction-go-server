package main

import (
	"log"
	"strconv"

	"github.com/Just-A-NoobieDev/auction-go-server/config"
	"github.com/Just-A-NoobieDev/auction-go-server/internal/auction"
	"github.com/Just-A-NoobieDev/auction-go-server/internal/bidding"
	"github.com/Just-A-NoobieDev/auction-go-server/internal/user"
	"github.com/Just-A-NoobieDev/auction-go-server/pkg/websocket"
	"github.com/Just-A-NoobieDev/auction-go-server/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := Init()

	// Inject dependencies
	handlers := Inject(db)


	router := router.NewRouter()
	router.SetupRouter(handlers)

	go websocket.HandleMessages()
	router.Engine.GET("/ws", func(c *gin.Context) {
		websocket.HandleConnections(c.Writer, c.Request)
	})

	log.Fatal(router.Engine.Run(":" + strconv.Itoa(config.AppConfig.Server.Port)))
}


func Init() *gorm.DB {
	config.Load()
	db := config.ConnectDatabase()

	return db
}

func Inject(db *gorm.DB) *router.Handlers {

	// Inject user dependencies
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	// Inject auction dependencies
	auctionRepository := auction.NewAuctionRepository(db)
	auctionService := auction.NewAuctionService(auctionRepository, userRepository)
	auctionHandler := auction.NewAuctionHandler(auctionService)

	// Inject bid dependencies
	bidRepository := bidding.NewBidRepository(db)
	bidService := bidding.NewBidService(bidRepository, auctionRepository)
	bidHandler := bidding.NewBidHandler(bidService)
	

	return &router.Handlers{
		UserHandler: userHandler,
		AuctionHandler: auctionHandler,
		BidHandler: bidHandler,
	}
}




