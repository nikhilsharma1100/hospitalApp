package doctor

import (
	"hospitalApp/initializers"
	"hospitalApp/internal/patient"
	"log"
)

func GetAll() []Doctor {
	var doctor []Doctor
	//context.JSON(http.StatusOK, gin.H{"data": initializers.Database.Find(&doctor)})
	initializers.Database.Preload("Patients").Find(&doctor)

	return doctor
}

func GetDoctorPatients(name string) ([]patient.Patient, error) {
	var doctor Doctor
	err := initializers.Database.Preload("Patients").Where(&Doctor{Name: name}).Find(&doctor).Error

	if err != nil {
		return []patient.Patient{}, err
	}

	return doctor.Patients, nil
}

func FindUserById(id uint) (Doctor, error) {
	var doctor Doctor
	err := initializers.Database.Where(&Doctor{DoctorId: id}).Find(&doctor).Error
	if err != nil {
		return Doctor{}, err
	}
	return doctor, nil
}

func FindUserByName(name string) (Doctor, error) {
	var doctor Doctor
	result := initializers.Database.Where(&Doctor{Name: name}).Find(&doctor)
	if result.Error != nil || result.RowsAffected == 0 {
		return Doctor{}, result.Error
	}
	return doctor, nil
}

func Create(entity Doctor) {
	log.Println("Before create:")
	result := initializers.Database.Omit("Patients").Create(&entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("After create:")
	log.Println(result.RowsAffected)
}

func Update(entity Doctor) {
	//doctor := entity
	//fmt.Println(entity)
	initializers.Database.Save(&entity)

	//fmt.Println(entity)
}

func UpdateAssociation(doctorEntity Doctor, entity patient.Patient) {
	err := initializers.Database.Model(&doctorEntity).Association("Patients").Append(&entity)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(entity Doctor) {
	result := initializers.Database.Delete(entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.RowsAffected)
}

func DeleteDoctorPatientRecord(name string) {
	//var patients patient.Patient
	//initializers.Database.Model(&Doctor{}).Where(Doctor{Name: name}).Association("Patients").Find(&patients)
	//log.Println(patients)

	//TODO : This is not working
	initializers.Database.Association("Patients").Delete(patient.Patient{Name: name})
}
