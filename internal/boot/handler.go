package boot

import (
	"context"
	"github.com/gin-gonic/gin"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/patient"
)

type Handler struct {
	doctorServer  *doctor.Server
	patientServer *patient.Server
	Engine        *gin.Engine
}

func (h *Handler) NewHandler() *Handler {
	return &Handler{
		doctorServer:  doctor.NewServer(),
		patientServer: patient.NewServer(),
		Engine:        h.ServeRoutes(gin.Default()),
	}
}

func (h *Handler) ServeRoutes(server *gin.Engine) *gin.Engine {
	router := server.Group("api")

	router = h.doctorRoutes(router)
	router = h.patientRoutes(router)

	return server
}

type handler struct {
	ctx context.Context
	Gin *gin.Engine
}

func (h *Handler) doctorRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("doctor", h.doctorServer.Core.GetAll) // call by server.go
	//router.GET("doctor/:id", h.doctorServer.Core.GetById)
	router.GET("doctor/:name", h.doctorServer.Core.GetByName)
	router.GET("doctor/getPatients", h.doctorServer.Core.GetPatient)
	router.POST("doctor", h.doctorServer.Core.Create)
	router.PATCH("doctor/:id", h.doctorServer.Core.Update)
	router.PATCH("doctor/addPatient", h.doctorServer.Core.UpdatePatientById)
	router.GET("doctor/deletePatient", h.doctorServer.Core.DeletePatient)

	return router
}

func (h *Handler) patientRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("patient", h.patientServer.Core.GetAll)
	router.GET("patient/:name", h.patientServer.Core.GetByName)
	router.POST("patient", h.patientServer.Core.Create)
	router.PATCH("patient/:id", h.patientServer.Core.Update)

	return router
}
