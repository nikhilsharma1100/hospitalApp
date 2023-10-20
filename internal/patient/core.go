package patient

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type ICore interface {
	GetByName(context *gin.Context)
	GetAll(context *gin.Context)
	GetById(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
}

func GetByName(context *gin.Context) {
	uri := UpdatePatientRequestUriName{}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	patient, err := GetEntityByName(uri.Name)
	log.Printf("Patient data get by Name(%q) : %+v", uri.Name, patient)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if patient.ID == "" {
		context.JSON(http.StatusOK, gin.H{"data": ""})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": patient})
}

func GetAll(context *gin.Context) {
	patients := GetAllEntities()

	context.JSON(http.StatusOK, gin.H{"data": patients})
}

func GetById(context *gin.Context) {
	uri := UpdatePatientRequestUri{}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	patient, err := GetEntityById(uri.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": patient})
}

func Create(context *gin.Context) {
	// Read request input here
	var inputData CreatePatientRequest
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var patientData Patient
	patientData.ID = generatePrimaryKey(5)
	patientData.Name = inputData.Name
	patientData.ContactNo = inputData.ContactNo
	patientData.Address = inputData.Address
	patientData.DoctorID = inputData.DoctorID
	patientData.CreatedAt = time.Now()
	patientData.UpdatedAt = time.Now()

	log.Printf("Patient data : %+v", patientData)
	CreateEntity(patientData)

	context.JSON(http.StatusCreated, gin.H{"data": "created"})
}

func Update(context *gin.Context) {
	// Read request input here
	inputData := UpdatePatientRequest{}
	uri := UpdatePatientRequestUri{}
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Printf("Patient data input : %+v", inputData)
	patientData, err := getPatientFromDBById(uri.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Patient data getFromDB : %+v", patientData)
	patientData.ContactNo = inputData.ContactNo
	patientData.Address = inputData.Address
	patientData.UpdatedAt = time.Now()
	UpdateEntity(patientData)

	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}

func getPatientFromDBById(id string) (Patient, error) {
	patient, err := GetEntityById(id)
	if err != nil {
		return Patient{}, err
	}

	return patient, nil
}

func generatePrimaryKey(length uint) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
