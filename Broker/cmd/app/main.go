package main

import (
	"mr-l0n3lly/go-broker/pkg/socket"
)

func main() {
	socketServer := socket.GetSocketServer()

	socketServer.Start()
}
