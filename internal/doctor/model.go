package doctor

import (
	"time"
)

type Doctor struct {
	DoctorId  uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:50"`
	ContactNo string `gorm:"size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllEntities() []Doctor {
	doctors := GetAll()

	return doctors
}

func GetEntityById(id uint) (Doctor, error) {
	doctor, err := FindUserById(id)
	if err != nil {
		return Doctor{}, err
	}

	return doctor, nil
}

func CreateEntity(entity Doctor) {
	Create(entity)
}

func UpdateEntity(entity Doctor) {
	Update(entity)
}
