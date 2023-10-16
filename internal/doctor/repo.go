package doctor

import (
	"hospitalApp/initializers"
	"log"
)

func GetAll() []Doctor {
	doctor := []Doctor{}
	//context.JSON(http.StatusOK, gin.H{"data": initializers.Database.Find(&doctor)})
	initializers.Database.Find(&doctor)

	return doctor
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
	result := initializers.Database.Create(&entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.RowsAffected)
}

func Update(entity Doctor) {
	//doctor := entity
	//fmt.Println(entity)
	initializers.Database.Save(&entity)

	//fmt.Println(entity)
}
func Delete(entity Doctor) {
	result := initializers.Database.Delete(entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.RowsAffected)
}
