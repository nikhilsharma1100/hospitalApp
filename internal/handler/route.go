package handler

import (
	"github.com/gin-gonic/gin"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/patient"
)

func ServeRoutes(server *gin.Engine) *gin.Engine {
	router := server.Group("api")

	router = doctorRoutes(router)
	router = patientRoutes(router)

	return server
}

func doctorRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("doctor/getAll", doctor.GetAllEntities)
	router.GET("doctor/getByName", doctor.GetEntityByName)
	router.GET("doctor/getPatients", doctor.GetPatientsByDoctor)
	router.POST("doctor/add", doctor.CreateEntity)
	router.PATCH("doctor/update", doctor.UpdateEntity)
	router.PATCH("doctor/addPatient", doctor.UpdatePatientDataById)
	router.GET("doctor/deletePatient", doctor.DeletePatientRecord)

	return router
}

func patientRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("patient/getAll", patient.GetAllEntities)
	router.GET("patient/getByName", patient.GetEntityByName)
	router.POST("patient/add", patient.CreateEntity)
	router.PATCH("patient/update", patient.UpdateEntity)

	return router
}
