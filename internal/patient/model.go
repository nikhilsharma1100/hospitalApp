package patient

import (
	"time"
)

type Patient struct {
	PatientID uint   `json:"patient_id" gorm:"size:5; primary_key"`
	Name      string `json:"name" gorm:"size:50; primary_key"`
	ContactNo string `json:"contact_no" gorm:"size:10; not null"`
	Address   string `json:"address" gorm:"size:50"`
	//Doctors   []*doctor.Doctor `gorm:"many2many:doctor_patients"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
