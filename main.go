package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hospitalApp/initializers"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/patient"
	"log"
)

var (
	server *gin.Engine
)

func init() {
	loadEnv()
	loadDatabase()
}

func loadDatabase() {
	initializers.Connect()
	initializers.Database.AutoMigrate(&doctor.Doctor{})
	initializers.Database.AutoMigrate(&patient.Patient{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	//doctorObj := doctor.Doctor{DoctorId: 6, Name: "Superman", ContactNo: "9876543210", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	//doctor.Create(doctorObj)

	//doctor.Update(doctorObj)
	//doctor.Delete(doctorObj)
	fmt.Println(doctor.GetAll())
}
