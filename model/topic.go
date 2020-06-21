package model

import (
	"github.com/jinzhu/gorm"
)

// Topic of chats
type Topic struct {
	gorm.Model

	Order uint
	Name  string
	Code  string
}
