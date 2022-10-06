package main

import (
	"mr-l0n3lly/go-broker/internal/db"
	"mr-l0n3lly/go-broker/internal/models"
	"mr-l0n3lly/go-broker/pkg/logging"
	"mr-l0n3lly/go-broker/pkg/socket"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	logger := logging.GetLogger()
	d, err := gorm.Open(sqlite.Open("broker.db"), &gorm.Config{})

	if err != nil {
		logger.Panic("failed to connect to database")
	}

	d.AutoMigrate(&models.Topic{})
	d.AutoMigrate(&models.Sender{})

	database := db.Init(d)
	socketServer := socket.GetSocketServer()

	socketServer.Start(*database)
}
