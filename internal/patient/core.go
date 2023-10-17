package patient

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type CreateRequestPatientData struct {
	PatientId uint   `json:"patient_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
	Address   string `json:"address"`
}

type UpdateRequestPatientData struct {
	PatientId uint   `json:"patient_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
	Address   string `json:"address"`
}

func GetEntityByName(context *gin.Context) {
	name := context.Query("name")
	patient, err := FindUserByName(name)
	log.Printf("Patient data get by Name(%q) : %+v", name, patient)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if patient.PatientID == 0 {
		context.JSON(http.StatusOK, gin.H{"data": ""})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": patient})
}

func GetAllEntities(context *gin.Context) {
	patients := GetAll()

	context.JSON(http.StatusOK, gin.H{"data": patients})
}

func GetEntityById(context *gin.Context, id uint) {
	patient, err := FindUserById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": patient})
}

func CreateEntity(context *gin.Context) {
	// Read request input here
	var inputData CreateRequestPatientData
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var patientData Patient
	patientData.PatientID = inputData.PatientId
	patientData.Name = inputData.Name
	patientData.ContactNo = inputData.ContactNo
	patientData.Address = inputData.Address
	patientData.CreatedAt = time.Now()
	patientData.UpdatedAt = time.Now()

	log.Printf("Patient data : %+v", patientData)
	Create(patientData)

	context.JSON(http.StatusCreated, gin.H{"data": "created"})
}

func UpdateEntity(context *gin.Context) {
	// Read request input here
	var inputData UpdateRequestPatientData
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Printf("Patient data input : %+v", inputData)
	patientData, err := GetPatientFromDBById(inputData.PatientId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Patient data getFromDB : %+v", patientData)
	patientData.Name = inputData.Name
	patientData.ContactNo = inputData.ContactNo
	patientData.UpdatedAt = time.Now()
	Update(patientData)

	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}

func GetPatientFromDBById(id uint) (Patient, error) {
	patient, err := FindUserById(id)
	if err != nil {
		return Patient{}, err
	}

	return patient, nil
}
