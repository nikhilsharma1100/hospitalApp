package boot

import (
	"github.com/gin-gonic/gin"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/patient"
)

type handler struct {
	doctorServer  *doctor.Server
	patientServer *patient.Server
	Engine        *gin.Engine
}

func NewHandler() *handler {
	h := &handler{}
	return &handler{
		doctorServer:  doctor.NewServer(),
		patientServer: patient.NewServer(),
		Engine:        h.ServeRoutes(),
	}
}

func (h *handler) ServeRoutes() *gin.Engine {

	h.doctorRoutes()
	h.patientRoutes()

	return h.Engine
}

//type handler struct {
//	ctx context.Context
//	Gin *gin.Engine
//}

func (h *handler) doctorRoutes() *gin.RouterGroup {
	router := h.Engine.Group("api")
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

func (h *handler) patientRoutes() *gin.RouterGroup {
	router := h.Engine.Group("api")
	router.GET("patient", h.patientServer.Core.GetAll)
	router.GET("patient/:name", h.patientServer.Core.GetByName)
	router.POST("patient", h.patientServer.Core.Create)
	router.PATCH("patient/:id", h.patientServer.Core.Update)

	return router
}
