package patient

import (
	"gorm.io/gorm"
	"hospitalApp/initializers"
	"log"
)

type Repo struct {
	db *gorm.DB
}

type IRepo interface {
	GetAllEntities() []Patient
	GetEntityById(id string) (Patient, error)
	GetEntityByName(name string) (Patient, error)
	CreateEntity(entity Patient)
	UpdateEntity(entity Patient)
	DeleteEntity(entity Patient)
}

func NewRepo() *Repo {
	return &Repo{
		db: initializers.Database,
	}
}

func (r *Repo) GetAllEntities() []Patient {
	patient := []Patient{}
	//context.JSON(http.StatusOK, gin.H{"data": initializers.Database.Find(&patient)})
	r.db.Find(&patient)

	return patient
}

func (r *Repo) GetEntityById(id string) (Patient, error) {
	var patient Patient
	err := r.db.Where(&Patient{ID: id}).Find(&patient).Error
	if err != nil {
		return Patient{}, err
	}
	return patient, nil
}

func (r *Repo) GetEntityByName(name string) (Patient, error) {
	var patient Patient
	result := r.db.Where(&Patient{Name: name}).Find(&patient)
	if result.Error != nil || result.RowsAffected == 0 {
		return Patient{}, result.Error
	}
	return patient, nil
}

func (r *Repo) CreateEntity(entity Patient) {
	result := r.db.Create(&entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.RowsAffected)
}

func (r *Repo) UpdateEntity(entity Patient) {
	//patient := entity
	//fmt.Println(entity)
	r.db.Save(&entity)

	//fmt.Println(entity)
}
func (r *Repo) DeleteEntity(entity Patient) {
	result := r.db.Delete(entity)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	log.Println(result.RowsAffected)
}
