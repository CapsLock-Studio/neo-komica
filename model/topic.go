package model

// Topic of chats
type Topic struct {
	PublicModel

	Order uint   `gorm:"not null"`
	Name  string `gorm:"not null"`
	Code  string `gorm:"unique_index;not null"`
}
