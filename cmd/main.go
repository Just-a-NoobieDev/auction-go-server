package main

import (
	"log"

	"github.com/Just-A-NoobieDev/auction-go-server/config"
	"github.com/Just-A-NoobieDev/auction-go-server/internal/auction"
	"github.com/Just-A-NoobieDev/auction-go-server/internal/user"
	"github.com/Just-A-NoobieDev/auction-go-server/router"
	"gorm.io/gorm"
)

func main() {
	db := Init()

	// Inject dependencies
	handlers := Inject(db)

	router := router.NewRouter()
	router.SetupRouter(handlers)

	log.Fatal(router.Engine.Run(":8080"))
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
	

	return &router.Handlers{
		UserHandler: userHandler,
		AuctionHandler: auctionHandler,
	}
}




