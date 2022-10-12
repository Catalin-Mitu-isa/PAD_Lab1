package main

import (
	"mr-l0n3lly/go-broker/internal/db"
	"mr-l0n3lly/go-broker/internal/models"
	grpc_server "mr-l0n3lly/go-broker/pkg/grpc-server"
	"mr-l0n3lly/go-broker/pkg/logging"
	"mr-l0n3lly/go-broker/pkg/socket"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var waitGroup sync.WaitGroup

	logger := logging.GetLogger()
	d, err := gorm.Open(sqlite.Open("broker.db"), &gorm.Config{})

	if err != nil {
		logger.Panic("failed to connect to database")
	}

	if res := d.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
		logger.Panic("failed to enforce foreign_keys")
	}

	err = d.AutoMigrate(&models.Topic{})

	database := db.Init(d)
	socketServer := socket.GetSocketServer()
	grpcServer := grpc_server.GetGrpcServer()

	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		socketServer.Start(*database)
	}()

	go func() {
		defer waitGroup.Done()
		grpcServer.Start()
	}()

	waitGroup.Wait()
}
