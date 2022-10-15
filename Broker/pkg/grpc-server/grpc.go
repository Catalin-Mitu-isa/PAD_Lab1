package grpc_server

import (
	"google.golang.org/grpc"
	"mr-l0n3lly/go-broker/internal/config"
	"mr-l0n3lly/go-broker/pkg/logging"
	"net"
	"strconv"
)

type Server struct {
	GRPCServer *grpc.Server
	Logger     *logging.Logger
}

var instance *Server

func GetGrpcServer() *Server {
	if instance == nil {
		instance = &Server{
			GRPCServer: grpc.NewServer(),
			Logger:     logging.GetLogger(),
		}
	}

	return instance
}

func (s *Server) Start() {
	cfg := config.GetConfiguration()

	addr := net.TCPAddr{
		Port: cfg.GrpcServer.Port,
		IP:   net.ParseIP(cfg.GrpcServer.Host),
	}

	socket, err := net.ListenTCP("tcp", &addr)
	s.Logger.Info("started grpc server")
	s.Logger.Info(cfg.GrpcServer.Host + ":" + strconv.Itoa(cfg.GrpcServer.Port))
	if err != nil {
		s.Logger.Fatal("grpc listen port failed")
	}

	// s.grpcServer.RegisterService()

	if err = s.GRPCServer.Serve(socket); err != nil {
		s.Logger.Fatal("failed to serve: %v", err)
	}
}
