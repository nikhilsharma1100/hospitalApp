package main

import (
	"github.com/joho/godotenv"
	"hospitalApp/initializers"
	"hospitalApp/internal/boot"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/patient"
	"log"
)

func main() {
	loadEnv()
	initDatabase()
	// init handlers/routes.go
	serveApplication()
	//doctorObj := doctor.Doctor{ID: 6, Name: "Superman", ContactNo: "9876543210", CreatedAt: time.Now(), UpdatedAt: time.Now()}
}

func initDatabase() {
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
	server := boot.NewHandler()
	server.Engine.Run(":8000")
}
