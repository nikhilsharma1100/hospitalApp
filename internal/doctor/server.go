package doctor

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

func (s *Server) Get(context *gin.Context) {
	request := GetDoctorByIdRequest{}
	if err := context.BindUri(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor, err := s.core.GetById(request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": doctor})
	return
}

func (s *Server) Create(context *gin.Context) {
	var inputData CreateDoctorRequest
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctorData, err := s.core.Create(inputData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": doctorData})
	return
}

func (s *Server) Update(context *gin.Context) {
	requestBody := UpdateDoctorRequestBody{}
	uri := UpdateDoctorRequestUri{}
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var inputData UpdateDoctorRequest
	inputData.ContactNo = requestBody.ContactNo
	inputData.ID = uri.ID

	if err := s.core.Update(inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "updated"})
	return
}

func (s *Server) GetPatientByDoctorId(context *gin.Context) {
	inputData := GetPatientByDoctorIdRequest{}
	if err := context.BindUri(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patientsData, err := s.core.GetPatientByDoctorId(inputData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": patientsData})

}
