package patient

import (
	"hospitalApp/initializers"
	"log"
)

func GetAll() []Patient {
	patient := []Patient{}
	//context.JSON(http.StatusOK, gin.H{"data": initializers.Database.Find(&patient)})
	initializers.Database.Find(&patient)

	return patient
}

func FindUserById(id uint) (Patient, error) {
	var patient Patient
	err := initializers.Database.Where(&Patient{PatientID: id}).Find(&patient).Error
	if err != nil {
		return Patient{}, err
	}
	return patient, nil
}

func FindUserByName(name string) (Patient, error) {
	var patient Patient
	result := initializers.Database.Where(&Patient{Name: name}).Find(&patient)
	if result.Error != nil || result.RowsAffected == 0 {
		return Patient{}, result.Error
	}
	return patient, nil
}

func Create(entity Patient) {
	result := initializers.Database.Create(&entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.RowsAffected)
}

func Update(entity Patient) {
	//patient := entity
	//fmt.Println(entity)
	initializers.Database.Save(&entity)

	//fmt.Println(entity)
}
func Delete(entity Patient) {
	result := initializers.Database.Delete(entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.RowsAffected)
}
