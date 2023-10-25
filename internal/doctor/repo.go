package doctor

import (
	"hospitalApp/initializers"
	"hospitalApp/internal/patient"
	"log"
)

type IRepo interface {
	GetAllEntities() []Doctor
	GetPatientEntityByDoctorId(name string) ([]patient.Patient, error)
	GetEntityById(id string) (Doctor, error)
	GetEntityByName(name string) (Doctor, error)
	CreateEntity(entity Doctor)
	UpdateEntity(entity Doctor)
	UpdateEntityAssociation(doctorEntity Doctor, entity patient.Patient)
	DeleteEntity(entity Doctor)
	DeletePatientEntityForDoctor(name string)
}

func GetAllEntities() []Doctor {
	var doctor []Doctor
	//context.JSON(http.StatusOK, gin.H{"data": initializers.Database.Find(&doctor)})
	initializers.Database.Preload("Patients").Find(&doctor)

	return doctor
}

func GetPatientEntityByDoctorId(id string) ([]patient.Patient, error) {
	var doctor Doctor
	//var patients []patient.Patient
	err := initializers.Database.Preload("Patients").Where(&Doctor{ID: id}).Find(&doctor).Error
	//err := initializers.Database.Model(&doctor).Where("doctor_id = ?", id).Association("Patients").Find(&patients)
	//result := initializers.Database.Preload("Patients").Where("doctor_id = ?", id).Find(&doctor)
	if err != nil {
		return []patient.Patient{}, err
	}

	return doctor.Patients, nil
}

func GetEntityById(id string) (Doctor, error) {
	var doctor Doctor
	result := initializers.Database.Where(&Doctor{ID: id}).Find(&doctor)
	if result.Error != nil {
		return Doctor{}, result.Error
	}

	if result.RowsAffected == 0 {
		return Doctor{}, nil
	}
	return doctor, nil
}

func GetEntityByName(name string) (Doctor, error) {
	var doctor Doctor
	result := initializers.Database.Where(&Doctor{Name: name}).Find(&doctor)
	if result.Error != nil || result.RowsAffected == 0 {
		return Doctor{}, result.Error
	}
	return doctor, nil
}

func CreateEntity(entity Doctor) {
	log.Println("Before create:")
	result := initializers.Database.Omit("Patients").Create(&entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("After create:")
	log.Println(result.RowsAffected)
}

func UpdateEntity(entity Doctor) {
	//doctor := entity
	//fmt.Println(entity)
	initializers.Database.Save(&entity)

	//fmt.Println(entity)
}
