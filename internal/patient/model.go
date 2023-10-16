package patient

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	PatientID uuid.UUID `gorm:"size:5; primary_key"`
	Name      string    `gorm:"size:50; primary_key"`
	ContactNo string    `gorm:"size:10; not null"`
	Address   string    `gorm:"size:50"`
	DoctorId  uuid.UUID `gorm:"foreignKey:DoctorId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
