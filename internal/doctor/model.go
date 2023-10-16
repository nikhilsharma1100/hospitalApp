package doctor

import (
	"hospitalApp/internal/patient"
	"time"
)

type Doctor struct {
	DoctorId  uint            `json:"doctor_id" gorm:"primary_key; references"`
	Name      string          `json:"name" gorm:"size:50"`
	ContactNo string          `json:"contact_no" gorm:"size:10"`
	Patients  patient.Patient `gorm:"many2many"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
