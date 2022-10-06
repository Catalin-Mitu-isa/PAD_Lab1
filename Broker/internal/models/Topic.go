package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	TopicName string `gorm:"unique"`
}
