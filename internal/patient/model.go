package patient

import (
	"time"
)

type Patient struct {
	ID        string `gorm:"size:5; primary_key"`
	Name      string `gorm:"size:50"`
	ContactNo string `gorm:"size:10; not null"`
	Address   string `gorm:"size:50"`
	DoctorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
