package model

import (
	uuid "github.com/satori/go.uuid"
)

// MediaType - the Media type deinifition
type MediaType int

const (
	// MediaTypeImage - Post is image
	MediaTypeImage MediaType = 0
	// MediaTypeVideo - Post is video
	MediaTypeVideo MediaType = 1
)

// Media - the model of image/video
type Media struct {
	PublicModel

	Type MediaType
	// Image or video source URL
	RawSrc *string `gorm:"type:text"`

	UUID     uuid.UUID `gorm:"type:char(36);unique_index;not null"`
	FileType string
	Height   uint
	Width    uint

	Posts []Post `gorm:"many2many:post_media_map"`
}
