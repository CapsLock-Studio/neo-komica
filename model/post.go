package model

import (
	uuid "github.com/satori/go.uuid"
)

// Post under the topic
type Post struct {
	PublicModel

	UUID     uuid.UUID `gorm:"type:char(36);unique_index;not null"`
	TopicID  uint
	Topic    Topic
	ParentID *uint `json:"-"`
	Parent   *Post
	Message  *string `gorm:"type:mediumtext"`

	Medias []Media `gorm:"many2many:post_media_map"`
}
