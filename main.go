package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hospitalApp/initializers"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/patient"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
	//doctorObj := doctor.Doctor{DoctorId: 6, Name: "Superman", ContactNo: "9876543210", CreatedAt: time.Now(), UpdatedAt: time.Now()}
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

	router := server.Group("api")
	router.GET("doctor/getAll", doctor.GetAllEntities)
	router.GET("doctor/getByName", doctor.GetEntityByName)
	router.GET("doctor/getPatients", doctor.GetPatientsByDoctor)
	router.POST("doctor/add", doctor.CreateEntity)
	router.PATCH("doctor/update", doctor.UpdateEntity)
	router.PATCH("doctor/addPatient", doctor.UpdatePatientDataById)
	router.GET("doctor/deletePatient", doctor.DeletePatientRecord)

	router.GET("patient/getAll", patient.GetAllEntities)
	router.GET("patient/getByName", patient.GetEntityByName)
	router.POST("patient/add", patient.CreateEntity)
	router.PATCH("patient/update", patient.UpdateEntity)
	
	server.Run(":8000")
}
