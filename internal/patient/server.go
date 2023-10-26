package patient

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	core *Core
}

func NewServer(c *Core) *Server {
	return &Server{c}
}

func (s *Server) GetAll(context *gin.Context) {
	doctors := s.core.GetAll()

	context.JSON(http.StatusOK, gin.H{"data": doctors})
	return
}

func (s *Server) Create(context *gin.Context) {
	// Read request input here
	var inputData CreatePatientRequest
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patientData, err := s.core.Create(inputData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": patientData})
	return
}

func (s *Server) Update(context *gin.Context) {
	// Read request input here
	requestBody := UpdatePatientRequestBody{}
	uri := UpdatePatientRequestUri{}
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inputData UpdatePatientRequest
	inputData.ContactNo = requestBody.ContactNo
	inputData.ID = uri.ID

	if err := s.core.Update(inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "updated"})
	return
}
