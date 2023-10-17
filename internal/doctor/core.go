package doctor

import (
	"github.com/gin-gonic/gin"
	"hospitalApp/internal/patient"
	"log"
	"net/http"
	"time"
)

type CreateRequestDoctorData struct {
	DoctorId  uint   `json:"doctor_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
}

type UpdateRequestDoctorData struct {
	DoctorId  uint   `json:"doctor_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
}

type UpdateRequestPatientData struct {
	DoctorId uint            `json:"doctor_id"`
	Patient  patient.Patient `json:"patient"`
}

func GetEntityByName(context *gin.Context) {
	name := context.Query("name")
	doctor, err := FindUserByName(name)
	log.Printf("Doctor data get by Name(%q) : %+v", name, doctor)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if doctor.DoctorId == 0 {
		context.JSON(http.StatusOK, gin.H{"data": ""})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": doctor})
}

func GetAllEntities(context *gin.Context) {
	doctors := GetAll()

	context.JSON(http.StatusOK, gin.H{"data": doctors})
}

func GetPatientsByDoctor(context *gin.Context) {
	name := context.Query("name")

	patientsData, err := GetDoctorPatients(name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": patientsData})
}

func DeletePatientRecord(context *gin.Context) {
	name := context.Query("name")

	DeleteDoctorPatientRecord(name)

	context.JSON(http.StatusOK, gin.H{"data": "deleted"})
}

func GetEntityById(context *gin.Context, id uint) {
	doctor, err := FindUserById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": doctor})
}

func CreateEntity(context *gin.Context) {
	// Read request input here
	var inputData CreateRequestDoctorData
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var doctorData Doctor
	doctorData.DoctorId = inputData.DoctorId
	doctorData.Name = inputData.Name
	doctorData.ContactNo = inputData.ContactNo
	doctorData.CreatedAt = time.Now()
	doctorData.UpdatedAt = time.Now()

	log.Println("Doctor data : %+v", doctorData)
	Create(doctorData)

	context.JSON(http.StatusCreated, gin.H{"data": "created"})
}

func UpdateEntity(context *gin.Context) {
	// Read request input here
	var inputData UpdateRequestDoctorData
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Printf("Doctor data input : %+v", inputData)
	doctorData, err := GetDoctorFromDBById(inputData.DoctorId)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Doctor data getFromDB : %+v", doctorData)
	doctorData.Name = inputData.Name
	doctorData.ContactNo = inputData.ContactNo
	doctorData.UpdatedAt = time.Now()
	Update(doctorData)

	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}

func UpdatePatientDataById(context *gin.Context) {
	// Read request input here
	var inputData UpdateRequestPatientData
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Printf("Doctor data input : %+v", inputData)
	doctorData, err := GetDoctorFromDBById(inputData.DoctorId)
	if err != nil {
		log.Fatal(err)
	}

	UpdateAssociation(doctorData, inputData.Patient)
	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}

func GetDoctorFromDBById(id uint) (Doctor, error) {
	doctor, err := FindUserById(id)
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}
