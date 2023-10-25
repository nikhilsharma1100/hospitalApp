package doctor

import (
	"hospitalApp/internal/patient"
	"time"
)

type Doctor struct {
	ID        string            `gorm:"primary_key; references; size:5"`
	Name      string            `gorm:"size:50"`
	ContactNo string            `gorm:"size:10"`
	Patients  []patient.Patient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
