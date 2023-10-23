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

// Core struct of account code
type Core struct {
	repo IRepo
}

type ICore interface {
	GetByName(context *gin.Context)
	GetAll(context *gin.Context)
	GetPatient(context *gin.Context)
	DeletePatient(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	UpdatePatientById(context *gin.Context)
}

func NewCore() *Core {
	return &Core{
		repo: NewRepo(),
	}
}

func (c *Core) GetByName(context *gin.Context) {
	uri := GetDoctorByNameRequest{}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor, err := c.repo.GetEntityByName(uri.Name)
	log.Printf("Doctor data get by Name(%q) : %+v", uri.Name, doctor)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if doctor.ID == "" {
		context.JSON(http.StatusOK, gin.H{"data": ""})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": doctor})
}

func (c *Core) GetAll(context *gin.Context) {
	doctors := c.repo.GetAllEntities()

	context.JSON(http.StatusOK, gin.H{"data": doctors})
}

func (c *Core) GetPatient(context *gin.Context) {
	name := context.Query("name")

	patientsData, err := c.repo.GetPatientEntityByName(name)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": patientsData})
}

func (c *Core) DeletePatient(context *gin.Context) {
	name := context.Query("name")

	c.repo.DeletePatientEntityForDoctor(name)

	context.JSON(http.StatusOK, gin.H{"data": "deleted"})
}

func (c *Core) Create(context *gin.Context) {
	// Read request input here
	var inputData CreateDoctorRequest
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Separate the logic and call from server.go
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
	c.repo.CreateEntity(doctorData)

	context.JSON(http.StatusCreated, gin.H{"data": doctorData})
}

func (c *Core) Update(context *gin.Context) {
	// Read request input here
	inputData := UpdateDoctorRequest{}
	uri := UpdateDoctorRequestUri{}
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := context.BindUri(&uri); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Doctor data input : %+v", inputData)

	doctorData, err := c.getDoctorFromDBById(uri.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Doctor data getFromDB : %+v", doctorData)
	doctorData.ContactNo = inputData.ContactNo
	doctorData.UpdatedAt = time.Now()
	c.repo.UpdateEntity(doctorData)

	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}

// Id in Name is confusing
func (c *Core) UpdatePatientById(context *gin.Context) {
	// Read request input here
	var inputData UpdatePatientRequest
	if err := context.ShouldBindJSON(&inputData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var emptyDoctorStruct Doctor
	doctorData, err := c.getDoctorFromDBById(inputData.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if reflect.DeepEqual(emptyDoctorStruct, doctorData) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Entity Id not found"})
		return
	}

	log.Println("Doctor data getById : %+v", doctorData)
	c.repo.UpdateEntityAssociation(doctorData, inputData.Patient)
	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}

func (c *Core) getDoctorFromDBById(id string) (Doctor, error) {
	doctor, err := c.repo.GetEntityById(id)
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
