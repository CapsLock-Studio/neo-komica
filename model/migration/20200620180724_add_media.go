package main

import (
	"github.com/CapsLock-Studio/neo-komica/model"
	"github.com/jinzhu/gorm"
)

type migration string

// Up - Changes for the migration.
func (m migration) Up(db *gorm.DB) error {
	db.AutoMigrate(&model.Media{})

	return nil
}

// Down - Rollback changes for the migration.
func (m migration) Down(db *gorm.DB) error {

	db.DropTableIfExists(&model.Media{})
	return nil
}

var Migration migration
