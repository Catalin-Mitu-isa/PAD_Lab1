package db

import (
	"mr-l0n3lly/go-broker/internal/models"

	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func Init(db *gorm.DB) *Database {
	return &Database{
		DB: db,
	}
}

func (db Database) GetSenders() ([]models.Sender, error) {
	senders := []models.Sender{}
	err := db.DB.Find(&senders).Error

	return senders, err
}

func (db Database) GetSenderById(id uint) (models.Sender, error) {
	sender := models.Sender{}
	err := db.DB.First(&sender, id).Error

	return sender, err
}

func (db Database) AddSender(s models.Sender) (uint, error) {
	err := db.DB.Create(&s).Error

	return s.ID, err
}

func (db Database) AddTopic(s models.Topic) (uint, error) {
	err := db.DB.Create(&s).Error

	return s.ID, err
}

func (db Database) GetTopics() ([]models.Topic, error) {
	topics := []models.Topic{}
	err := db.DB.Find(&topics).Error

	return topics, err
}

func (db Database) GetTopicById(id uint) (models.Topic, error) {
	topic := models.Topic{}
	err := db.DB.First(&topic, id).Error

	return topic, err
}
