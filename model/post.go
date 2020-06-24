package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Post under the topic
type Post struct {
	PublicModel

	UUID     uuid.UUID `gorm:"type:char(36);unique_index;not null"`
	TopicID  uint      `json:"-"`
	Topic    Topic     `json:"-"`
	ParentID *uint     `json:"-"`
	Parent   *Post
	Message  *string `gorm:"type:mediumtext"`

	Medias  []Media `gorm:"many2many:post_media_map"`
	Replies []Post  `gorm:"foreignkey:parent_id"`
}

// PostOrder - scope for post order.
func PostOrder(db *gorm.DB) *gorm.DB {
	return db.Order("id DESC")
}

// ReplyOrder - scope for reply order.
func ReplyOrder(db *gorm.DB) *gorm.DB {
	return db.Order("id ASC")
}
