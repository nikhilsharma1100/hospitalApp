package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primary_key"`
	Name      string    `gorm:"primary_key"`
	Email     string    `gorm:"type:varchar(255); not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
