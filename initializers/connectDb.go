package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s post=%s sslmode=disable TimeZone=Asia/Kolkata, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	fmt.Println("Connected successfully to the database")
}
