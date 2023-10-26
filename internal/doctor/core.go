package doctor

import (
	"hospitalApp/internal/patient"
	"log"
	"math/rand"
	"time"
)

type Core struct {
	repo *Repo
}

type ICore interface {
	GetById(request GetDoctorByIdRequest) (Doctor, error)
	GetAll() []Doctor
	GetPatientByDoctorId(request GetPatientByDoctorIdRequest) ([]patient.Patient, error)
	Create(request CreateDoctorRequest) (Doctor, error)
	Update(request UpdateDoctorRequest) error
}

type IValidator interface {
	ValidateCreateRequest(input CreateDoctorRequest) error
}

func NewCore(r *Repo) *Core {
	return &Core{r}
}

func (c *Core) GetById(request GetDoctorByIdRequest) (Doctor, error) {

	doctor, err := c.repo.GetEntityById(request.ID)
	log.Printf("Doctor data get by Id(%q) : %+v", request.ID, doctor)
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func (c *Core) GetAll() []Doctor {

	doctors := c.repo.GetAllEntities()

	return doctors
}

func (c *Core) GetPatientByDoctorId(request GetPatientByDoctorIdRequest) ([]patient.Patient, error) {

	patientsData, err := c.repo.GetPatientEntityByDoctorId(request.ID)
	if err != nil {
		return []patient.Patient{}, err
	}

	return patientsData, nil
}

func (c *Core) Create(request CreateDoctorRequest) (Doctor, error) {

	validationErr := c.ValidateCreateRequest(request)
	if validationErr != nil {
		return Doctor{}, validationErr
	}

	var doctorData Doctor
	doctorData.ID = c.generatePrimaryKey(5)
	doctorData.Name = request.Name
	doctorData.ContactNo = request.ContactNo
	doctorData.CreatedAt = time.Now()
	doctorData.UpdatedAt = time.Now()

	log.Println("Doctor data : %+v", doctorData)
	c.repo.CreateEntity(doctorData)

	return doctorData, nil
}

func (c *Core) Update(request UpdateDoctorRequest) error {

	doctorData, err := c.getDoctorFromDBById(request.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Doctor data getFromDB : %+v", doctorData)

	doctorData.ContactNo = request.ContactNo
	doctorData.UpdatedAt = time.Now()
	c.repo.UpdateEntity(doctorData)

	return nil
}

func (c *Core) getDoctorFromDBById(id string) (Doctor, error) {

	doctor, err := c.repo.GetEntityById(id)
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
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
