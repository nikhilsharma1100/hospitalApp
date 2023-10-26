package doctor

import (
	"gorm.io/gorm"
	"hospitalApp/internal/patient"
	"log"
)

type Repo struct {
	db *gorm.DB
}

type IRepo interface {
	GetAllEntities() []Doctor
	GetPatientEntityByDoctorId(name string) ([]patient.Patient, error)
	GetEntityById(id string) (Doctor, error)
	CreateEntity(entity Doctor)
	UpdateEntity(entity Doctor)
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) GetAllEntities() []Doctor {
	var doctor []Doctor
	//context.JSON(http.StatusOK, gin.H{"data": initializers.Database.Find(&doctor)})
	r.db.Preload("Patients").Find(&doctor)

	return doctor
}

func (r *Repo) GetPatientEntityByDoctorId(id string) ([]patient.Patient, error) {
	var doctor Doctor
	//var patients []patient.Patient
	err := r.db.Where(&Doctor{ID: id}).Preload("Patients").Find(&doctor).Error
	//err := r.db.Model(&doctor).Where("doctor_id = ?", id).Association("Patients").Find(&patients)
	//result := r.db.Preload("Patients").Where("doctor_id = ?", id).Find(&doctor)
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
