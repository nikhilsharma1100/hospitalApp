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
	router.GET("doctor", doctor.GetAll) // call by server.go
	//router.GET("doctor/:id", doctor.GetById)
	router.GET("doctor/:name", doctor.GetByName)
	router.GET("doctor/getPatients", doctor.GetPatient)
	router.POST("doctor", doctor.Create)
	router.PATCH("doctor/:id", doctor.Update)
	router.PATCH("doctor/addPatient", doctor.UpdatePatientById)
	router.GET("doctor/deletePatient", doctor.DeletePatient)

	return router
}

func patientRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("patient", patient.GetAll)
	router.GET("patient/:name", patient.GetByName)
	router.POST("patient", patient.Create)
	router.PATCH("patient/:id", patient.Update)

	return router
}
