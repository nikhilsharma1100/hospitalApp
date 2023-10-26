package patient

import (
	"log"
	"math/rand"
	"time"
)

type Core struct {
	repo *Repo
}

type ICore interface {
	GetAll() []Patient
	Create(request CreatePatientRequest) (Patient, error)
	Update(request UpdatePatientRequest) error
}

type IValidator interface {
	ValidateCreateRequest(input CreatePatientRequest) error
}

func NewCore(r *Repo) *Core {
	return &Core{r}
}

func (c *Core) GetAll() []Patient {
	patients := c.repo.GetAllEntities()

	return patients
}

func (c *Core) Create(request CreatePatientRequest) (Patient, error) {

	validationErr := c.ValidateCreateRequest(request)
	if validationErr != nil {
		return Patient{}, validationErr
	}

	var patientData Patient
	patientData.ID = c.generatePrimaryKey(5)
	patientData.Name = request.Name
	patientData.ContactNo = request.ContactNo
	patientData.Address = request.Address
	patientData.DoctorID = request.DoctorID
	patientData.CreatedAt = time.Now()
	patientData.UpdatedAt = time.Now()

	log.Printf("Patient data : %+v", patientData)
	c.repo.CreateEntity(patientData)

	return patientData, nil
}

func (c *Core) Update(request UpdatePatientRequest) error {

	patientData, err := c.getPatientFromDBById(request.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Patient data getFromDB : %+v", patientData)
	patientData.DoctorID = request.DoctorID
	patientData.ContactNo = request.ContactNo
	patientData.Address = request.Address
	patientData.UpdatedAt = time.Now()
	c.repo.UpdateEntity(patientData)

	return nil
}

func (c *Core) getPatientFromDBById(id string) (Patient, error) {
	patient, err := c.repo.GetEntityById(id)
	if err != nil {
		return Patient{}, err
	}

	return patient, nil
}

func (c *Core) generatePrimaryKey(length uint) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
