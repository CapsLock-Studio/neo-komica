package model

import (
	"time"
)

// PublicModel base model definition but not output to JSON
type PublicModel struct {
	ID        uint `gorm:"primary_key" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
