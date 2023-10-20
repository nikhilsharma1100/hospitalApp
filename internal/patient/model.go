package patient

import (
	"time"
)

type Patient struct {
	ID        string `json:"id" gorm:"size:5; primary_key"`
	Name      string `json:"name" gorm:"size:50; primary_key"`
	ContactNo string `json:"contact_no" gorm:"size:10; not null"`
	Address   string `json:"address" gorm:"size:50"`
	DoctorID  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
