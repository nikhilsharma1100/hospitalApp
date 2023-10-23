package doctor

import (
	"gorm.io/gorm"
	"hospitalApp/initializers"
	"hospitalApp/internal/patient"
	"log"
)

type Repo struct {
	db *gorm.DB
}

type IRepo interface {
	GetAllEntities() []Doctor
	GetPatientEntityByName(name string) ([]patient.Patient, error)
	GetEntityById(id string) (Doctor, error)
	GetEntityByName(name string) (Doctor, error)
	CreateEntity(entity Doctor)
	UpdateEntity(entity Doctor)
	UpdateEntityAssociation(doctorEntity Doctor, entity patient.Patient)
	DeleteEntity(entity Doctor)
	DeletePatientEntityForDoctor(name string)
}

func NewRepo() *Repo {
	return &Repo{
		db: initializers.Database,
	}
}

func (r *Repo) GetAllEntities() []Doctor {
	var doctor []Doctor
	r.db.Preload("Patients").Find(&doctor)

	return doctor
}

func (r *Repo) GetPatientEntityByName(name string) ([]patient.Patient, error) {
	var doctor Doctor
	err := r.db.Preload("Patients").Where(&Doctor{Name: name}).Find(&doctor).Error

	if err != nil {
		return []patient.Patient{}, err
	}

	return doctor.Patients, nil
}

func (r *Repo) GetEntityById(id string) (Doctor, error) {
	var doctor Doctor
	result := r.db.Where(&Doctor{ID: id}).Find(&doctor)
	if result.Error != nil {
		return Doctor{}, result.Error
	}

	if result.RowsAffected == 0 {
		return Doctor{}, nil
	}
	return doctor, nil
}

func (r *Repo) GetEntityByName(name string) (Doctor, error) {
	var doctor Doctor
	result := r.db.Where(&Doctor{Name: name}).Find(&doctor)
	if result.Error != nil || result.RowsAffected == 0 {
		return Doctor{}, result.Error
	}
	return doctor, nil
}

func (r *Repo) CreateEntity(entity Doctor) {
	log.Println("Before create:")
	result := r.db.Omit("Patients").Create(&entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("After create:")
	log.Println(result.RowsAffected)
}

func (r *Repo) UpdateEntity(entity Doctor) {
	//doctor := entity
	//fmt.Println(entity)
	r.db.Save(&entity)

	//fmt.Println(entity)
}

func (r *Repo) UpdateEntityAssociation(doctorEntity Doctor, entity patient.Patient) {
	err := r.db.Model(&doctorEntity).Association("Patients").Append(&entity)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Repo) DeleteEntity(entity Doctor) {
	result := r.db.Delete(entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.RowsAffected)
}

func (r *Repo) DeletePatientEntityForDoctor(name string) {
	//var patients patient.Patient
	//r.db.Model(&Doctor{}).Where(Doctor{Name: name}).Association("Patients").Find(&patients)
	//log.Println(patients)

	//TODO : This is not working
	r.db.Association("Patients").Delete(patient.Patient{Name: name})
}
