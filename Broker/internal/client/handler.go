package client

import (
	"encoding/json"
	"fmt"
	"mr-l0n3lly/go-broker/internal/db"
	"mr-l0n3lly/go-broker/internal/messages"
	"mr-l0n3lly/go-broker/internal/models"
	"mr-l0n3lly/go-broker/pkg/logging"
	"net"
)

type Handler struct {
	DB db.Database
}

func (c Handler) createTopicAction(jsonData messages.SenderRequest) (response messages.SenderResponse) {
	// Add topic to database
	_, err := c.DB.AddTopic(models.Topic{
		TopicName: jsonData.TopicName,
	})

	if err != nil {
		response = messages.SenderResponse{
			Action:  messages.CreateTopicAction,
			Success: false,
			Error:   err.Error(),
		}
	} else {
		response = messages.SenderResponse{
			Action:  messages.CreateTopicAction,
			Success: true,
		}
	}

	return response
}

func (c Handler) publishMessage(jsonData messages.SenderRequest) (response messages.SenderResponse) {
	logger := logging.GetLogger()
	subscribers := models.Subscriber{}.GetTopicSubscribers(jsonData.TopicName)
	tmp := make([]models.Subscriber, 0)

	fmt.Println(subscribers)
	for index, sub := range subscribers {
		logger.Info("sending message", jsonData.Message)
		_, err := (*sub.Conn).Write([]byte(jsonData.Message))
		if err != nil {
			logger.Error(err)
			(*sub.Conn).Close()
			tmp = append(subscribers[:index], subscribers[index+1])
		}
	}

	subscribers = tmp

	return messages.SenderResponse{
		Action:  jsonData.Action,
		Success: true,
	}
}

func (c Handler) subscribe(jsonData messages.SenderRequest, conn *net.Conn) messages.SenderResponse {
	models.Subscriber{}.AddSubscriber(jsonData.TopicName, conn)

	return messages.SenderResponse{
		Action:  jsonData.Action,
		Success: true,
	}
}

func (c Handler) ParseMessage(jsonData messages.SenderRequest, conn *net.Conn) ([]byte, error) {
	var responseJson []byte
	var response messages.SenderResponse
	var err error

	switch jsonData.Action {
	case messages.CreateTopicAction:
		// Craft a response for sender
		response = c.createTopicAction(jsonData)

	case messages.PublishMessageAction:
		// Publish a message
		response = c.publishMessage(jsonData)

	case messages.SubscribeAction:
		// Subscribe
		response = c.subscribe(jsonData, conn)
	}

	responseJson, _ = json.Marshal(response)

	return responseJson, err
}
