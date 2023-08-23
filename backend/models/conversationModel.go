package models

import (
	"gorm.io/gorm"
)

type Conversation struct {
	gorm.Model
	Messages []Message `json:"messages"`
}
