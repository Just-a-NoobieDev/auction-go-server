package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func ConnectDatabase() *gorm.DB {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
        AppConfig.Database.Host,
        AppConfig.Database.User,
        AppConfig.Database.Password,
        AppConfig.Database.DBName,
        AppConfig.Database.Port,
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    return db
}