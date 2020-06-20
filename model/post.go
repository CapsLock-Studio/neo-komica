package model

import (
	"github.com/jinzhu/gorm"
)

// Post under the topic
type Post struct {
	gorm.Model

	TopicID  uint
	ParentID *uint
	Parent   *Post
	Message  *string `gorm:"type:mediumtext"`

	Medias []Media `gorm:"many2many:post_media_map"`
}
