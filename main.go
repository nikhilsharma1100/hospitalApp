package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hospitalApp/initializers"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/handler"
	"hospitalApp/internal/patient"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
	//doctorObj := doctor.Doctor{ID: 6, Name: "Superman", ContactNo: "9876543210", CreatedAt: time.Now(), UpdatedAt: time.Now()}
}

func loadDatabase() {
	initializers.Connect()
	initializers.Database.AutoMigrate(&doctor.Doctor{})
	initializers.Database.AutoMigrate(&patient.Patient{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env.local file")
	}
}

func serveApplication() {
	server := gin.Default()

	server = handler.ServeRoutes(server)
	server.Run(":8000")
}
