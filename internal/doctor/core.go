package doctor

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"time"
)

type ICore interface {
	GetByName(context *gin.Context)
	GetAll(context *gin.Context)
	GetPatient(context *gin.Context)
	DeletePatient(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	UpdatePatientById(context *gin.Context)
}

func GetByName(context *gin.Context) {
	uri := GetDoctorByNameRequest{}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	doctor, err := GetEntityByName(uri.Name)
	log.Printf("Doctor data get by Name(%q) : %+v", uri.Name, doctor)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if doctor.ID == "" {
		context.JSON(http.StatusOK, gin.H{"data": ""})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": doctor})
}

func GetAll(context *gin.Context) {
	doctors := GetAllEntities()

	context.JSON(http.StatusOK, gin.H{"data": doctors})
}

func GetPatient(context *gin.Context) {
	name := context.Query("name")

	patientsData, err := GetPatientEntityByName(name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": patientsData})
}

func DeletePatient(context *gin.Context) {
	name := context.Query("name")

	DeletePatientEntityForDoctor(name)

	context.JSON(http.StatusOK, gin.H{"data": "deleted"})
}

func Create(context *gin.Context) {
	// Read request input here
	var inputData CreateDoctorRequest
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	validationErr := validation.ValidateStruct(&inputData,
		validation.Field(&inputData.ContactNo, validation.Match(regexp.MustCompile("\\d{10}$")), validation.Length(10, 10)),
	)

	if validationErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	var doctorData Doctor
	doctorData.ID = generatePrimaryKey(5)
	doctorData.Name = inputData.Name
	doctorData.ContactNo = inputData.ContactNo
	doctorData.CreatedAt = time.Now()
	doctorData.UpdatedAt = time.Now()

	log.Println("Doctor data : %+v", doctorData)
	CreateEntity(doctorData)

	context.JSON(http.StatusCreated, gin.H{"data": doctorData})
}

func Update(context *gin.Context) {
	// Read request input here
	inputData := UpdateDoctorRequest{}
	uri := UpdateDoctorRequestUri{}
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	log.Printf("Doctor data input : %+v", inputData)

	doctorData, err := getDoctorFromDBById(uri.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Doctor data getFromDB : %+v", doctorData)
	doctorData.ContactNo = inputData.ContactNo
	doctorData.UpdatedAt = time.Now()
	UpdateEntity(doctorData)

	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}

func UpdatePatientById(context *gin.Context) {
	// Read request input here
	var inputData UpdatePatientRequest
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var emptyDoctorStruct Doctor
	doctorData, err := getDoctorFromDBById(inputData.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if reflect.DeepEqual(emptyDoctorStruct, doctorData) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Entity Id not found"})
		return
	}

	log.Println("Doctor data getById : %+v", doctorData)
	UpdateEntityAssociation(doctorData, inputData.Patient)
	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}

func getDoctorFromDBById(id string) (Doctor, error) {
	doctor, err := GetEntityById(id)
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
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
