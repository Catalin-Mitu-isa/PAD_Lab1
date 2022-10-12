package grpc_server

import (
	"google.golang.org/grpc"
	"mr-l0n3lly/go-broker/internal/config"
	"mr-l0n3lly/go-broker/pkg/logging"
	"net"
	"strconv"
)

type Server struct {
	grpcServer *grpc.Server
}

var instance *Server

func GetGrpcServer() *Server {
	if instance == nil {
		instance = &Server{
			grpcServer: grpc.NewServer(),
		}
	}

	return instance
}

func (s *Server) Start() {
	cfg := config.GetConfiguration()
	logger := logging.GetLogger()

	addr := net.TCPAddr{
		Port: cfg.GrpcServer.Port,
		IP:   net.ParseIP(cfg.GrpcServer.Host),
	}

	socket, err := net.ListenTCP("tcp", &addr)
	logger.Info("started grpc server")
	logger.Info(cfg.GrpcServer.Host + ":" + strconv.Itoa(cfg.GrpcServer.Port))
	if err != nil {
		logger.Fatal("grpc listen port failed")
	}

	s.grpcServer.RegisterService()

	if err = s.grpcServer.Serve(socket); err != nil {
		logger.Fatal("failed to serve: %v", err)
	}
}
