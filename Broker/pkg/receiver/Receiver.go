package receiver

import (
	"context"
	"encoding/json"
	"mr-l0n3lly/go-broker/pkg/logging"
)

type GrpcSubscriber struct {
	TopicName string
	HostName  string
	Port      int32
}

var subscribers = make([]GrpcSubscriber, 0)

type ReceiverService struct {
	UnimplementedReceiverServiceServer
}

func (r *ReceiverService) Subscribe(ctx context.Context, request *SubscribeRequest) (*SubscribeResponse, error) {
	logger := logging.GetLogger()

	logger.Info("subscribing client")

	subscribers = append(subscribers, GrpcSubscriber{
		TopicName: request.TopicName,
		HostName:  request.HostName,
		Port:      request.Port,
	})

	response, _ := json.Marshal(SubscribeResponse{
		Success: true,
	})

	logger.Info(string(response))

	return &SubscribeResponse{
		Success: true,
	}, nil
}

func GetGRPCSubscribers() []GrpcSubscriber {
	return subscribers
}
