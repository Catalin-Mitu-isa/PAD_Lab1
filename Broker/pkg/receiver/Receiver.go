package receiver

import (
	"context"
	grpc_server "mr-l0n3lly/go-broker/pkg/grpc-server"
)

type Receiver struct {
}

func (*Receiver) Subscribe(ctx context.Context, request *grpc_server.SubscribeRequest) *grpc_server.SubscribeResponse {
	topicName := request.TopicName

	response := &grpc_server.SubscribeResponse{
		Success: true,
	}

	return response
}
