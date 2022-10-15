package sender

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"mr-l0n3lly/go-broker/internal/db"
	"mr-l0n3lly/go-broker/internal/models"
	"mr-l0n3lly/go-broker/pkg/broker"
	"mr-l0n3lly/go-broker/pkg/logging"
	"mr-l0n3lly/go-broker/pkg/receiver"
	"strconv"
)

type SenderService struct {
	UnimplementedSenderServiceServer
}

func (s *SenderService) CreateTopic(ctx context.Context, request *CreateTopicRequest) (*CreateTopicResponse, error) {
	logger := logging.GetLogger()
	database := db.GetDatabase()

	logger.Info("subscribing client")

	_, err := database.AddTopic(models.Topic{
		TopicName: request.Name,
	})

	if err != nil {
		return &CreateTopicResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	response, _ := json.Marshal(CreateTopicResponse{
		Success: true,
	})

	logger.Info(string(response))

	return &CreateTopicResponse{
		Success: true,
	}, nil
}

func (s *SenderService) PublishMessage(ctx context.Context, request *PublishMessageRequest) (*PublishMessageResponse, error) {
	topicName := request.TopicName
	subscribers := receiver.GetGRPCSubscribers()
	logger := logging.GetLogger()

	for _, sub := range subscribers {
		if sub.TopicName == topicName {
			conn, err := grpc.Dial(sub.HostName + ":" + strconv.Itoa(int(sub.Port)))
			if err != nil {
				continue
			}

			payload := &broker.SendMessageRequest{
				Message: request.Message,
			}

			client := broker.NewBrokerServiceClient(conn)

			_, err = client.SendMessage(context.Background(), payload)

			if err != nil {
				continue
			}
		}
	}

	response, _ := json.Marshal(PublishMessageResponse{
		Success: true,
	})
	logger.Info(string(response))

	return &PublishMessageResponse{
		Success: true,
	}, nil
}
