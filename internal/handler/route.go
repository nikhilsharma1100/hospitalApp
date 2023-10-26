package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/patient"
)

func ServeRoutes(server *gin.Engine, db *gorm.DB) (*gin.Engine, error) {
	router := server.Group("api")
	
	doctorRepo := doctor.NewRepo(db)
	patientRepo := patient.NewRepo(db)

	doctorCore := doctor.NewCore(doctorRepo)
	patientCore := patient.NewCore(patientRepo)

	doctorServer := doctor.NewServer(doctorCore)
	patientServer := patient.NewServer(patientCore)

	router = doctorRoutes(router, doctorServer)
	router = patientRoutes(router, patientServer)

	return server, nil
}

func doctorRoutes(router *gin.RouterGroup, server *doctor.Server) *gin.RouterGroup {
	router.POST("doctor", server.Create)
	router.PATCH("doctor/:id", server.Update)
	router.GET("doctor", server.GetAll)
	router.GET("doctor/:id", server.Get)
	router.GET("doctor/getPatientsByDoctorId/:id", server.GetPatientByDoctorId)

	return router
}

func patientRoutes(router *gin.RouterGroup, server *patient.Server) *gin.RouterGroup {
	router.POST("patient", server.Create)
	router.PATCH("patient/:id", server.Update)
	router.GET("patient", server.GetAll)

	return router
}
