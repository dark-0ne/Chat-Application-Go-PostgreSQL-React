package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Text           string `json:"text" gorm:"not null`
	Read           bool   `json:"read" gorm:"not null; default false"`
	SenderID       uint   `json:"sender_id" gorm:"not null"`
	ConversationID uint   `json:"conversation_id" gorm:"not null"`
}
