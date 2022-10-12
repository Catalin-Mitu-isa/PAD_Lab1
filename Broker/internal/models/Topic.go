package models

import (
	"time"
)

type Topic struct {
	ID        uint   `gorm:"primarykey"`
	TopicName string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
