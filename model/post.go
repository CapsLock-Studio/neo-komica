package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Post under the topic
type Post struct {
	gorm.Model

	UUID     uuid.UUID `gorm:"type:char(36);unique_index;not null"`
	TopicID  uint
	ParentID *uint
	Parent   *Post
	Message  *string `gorm:"type:mediumtext"`

	Medias []Media `gorm:"many2many:post_media_map"`
}
