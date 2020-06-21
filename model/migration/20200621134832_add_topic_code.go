package main

import (
	"github.com/jinzhu/gorm"

	"github.com/CapsLock-Studio/neo-komica/model"
)

type migration string

// Up - Changes for the migration.
func (m migration) Up(db *gorm.DB) error {
	db.AutoMigrate(&model.Topic{})
	return nil
}

// Down - Rollback changes for the migration.
func (m migration) Down(db *gorm.DB) error {
	return nil
}

var Migration migration
