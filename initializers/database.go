package initializers

import (
	"fmt"
	"hospitalApp/internal/doctor"
	"hospitalApp/internal/patient"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return Database, err
}

func RunMigrations(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&doctor.Doctor{})
	db.AutoMigrate(&patient.Patient{})

	return db
}
