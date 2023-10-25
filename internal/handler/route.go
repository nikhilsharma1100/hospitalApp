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
	router.GET("doctor", doctor.GetAll)
	//router.GET("doctor/:id", doctor.GetById)
	router.GET("doctor/:name", doctor.GetByName)
	router.GET("doctor/getPatientsByDoctorId/:id", doctor.GetPatientByDoctorId)
	router.POST("doctor", doctor.Create)
	router.PATCH("doctor/:id", doctor.Update)

	return router
}

func patientRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("patient", patient.GetAll)
	router.GET("patient/:name", patient.GetByName)
	router.POST("patient", patient.Create)
	router.PATCH("patient/:id", patient.Update)

	return router
}
