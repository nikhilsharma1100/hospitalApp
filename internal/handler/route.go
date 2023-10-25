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
	router.POST("doctor", doctor.Create)
	router.PATCH("doctor/:id", doctor.Update)
	router.GET("doctor", doctor.GetAll)
	router.GET("doctor/:id", doctor.GetById)
	router.GET("doctor/getPatientsByDoctorId/:id", doctor.GetPatientByDoctorId)

	return router
}

func patientRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.POST("patient", patient.Create)
	router.PATCH("patient/:id", patient.Update)
	router.GET("patient", patient.GetAll)

	return router
}
