package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string         `json:"username" gorm:"size:255;not null;unique"`
	Bio           string         `json:"bio" gorm:"size:255"`
	Password      string         `json:"password" gorm:"size:255;not null"`
	Conversations []Conversation `json:"conversations" gorm:"many2many:user_conversations"`
}
