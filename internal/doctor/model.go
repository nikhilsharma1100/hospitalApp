package doctor

import (
	"hospitalApp/internal/patient"
	"time"
)

type Doctor struct {
	ID        string `json:"id" gorm:"primary_key; references; size:5"`
	Name      string `json:"name" gorm:"size:50"`
	ContactNo string `json:"contact_no" gorm:"size:10"`
	Patients  []patient.Patient
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
