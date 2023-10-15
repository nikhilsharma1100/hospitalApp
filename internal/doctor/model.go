package model

import (
	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	// ID        uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primary_key"`
	Name      string `gorm:"primary_key"`
	ContactNo string `gorm:"type:varchar(10)"`
}
