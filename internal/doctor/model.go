package doctor

import (
	"time"
)

type Doctor struct {
	DoctorId  uint   `gorm:"primaryKey"`
	Name      string `json:"name" gorm:"size:50"`
	ContactNo string `json:"contact_no" gorm:"size:10"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
