package main

import (
	"github.com/jinzhu/gorm"

	"github.com/CapsLock-Studio/neo-komica/model"
)

type migration string

// Up - Changes for the migration.
func (m migration) Up(db *gorm.DB) error {
	db.AutoMigrate(&model.Post{})

	db.Model(&model.Post{}).AddForeignKey("topic_id", "topics(id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Post{}).AddForeignKey("parent_id", "posts(id)", "RESTRICT", "RESTRICT")
	return nil
}

// Down - Rollback changes for the migration.
func (m migration) Down(db *gorm.DB) error {
	db.Model(&model.Post{}).RemoveForeignKey("topic_id", "topics(id)")
	db.Model(&model.Post{}).RemoveForeignKey("parent_id", "posts(id)")

	db.DropTableIfExists(&model.Topic{})
	return nil
}

var Migration migration
