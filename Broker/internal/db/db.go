package db

import (
	"gorm.io/gorm"
	"mr-l0n3lly/go-broker/internal/models"
)

type Database struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) *Database {
	return &Database{
		DB: db,
	}
}

func (db Database) AddTopic(s models.Topic) (uint, error) {
	err := db.DB.Create(&s).Error

	return s.ID, err
}

func (db Database) GetTopics() ([]models.Topic, error) {
	var topics []models.Topic
	err := db.DB.Find(&topics).Error

	return topics, err
}

func (db Database) GetTopicById(id uint) (models.Topic, error) {
	topic := models.Topic{}
	err := db.DB.First(&topic, id).Error

	return topic, err
}
