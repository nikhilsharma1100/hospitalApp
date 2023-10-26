package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"hospitalApp/initializers"
	"hospitalApp/internal/handler"
	"log"
)

var db *gorm.DB

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
	//doctorObj := doctor.Doctor{ID: 6, Name: "Superman", ContactNo: "9876543210", CreatedAt: time.Now(), UpdatedAt: time.Now()}
}

func loadDatabase() {
	db, _ = initializers.Connect()

	log.Println("1")
	log.Println(db)
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}
}

func serveApplication() {
	server := gin.Default()

	log.Println("2")
	log.Println(db)
	server, err := handler.ServeRoutes(server, db)
	if err != nil {
		log.Fatal("Error while running the server: " + err.Error())
	}
	server.Run(":8000")
}
