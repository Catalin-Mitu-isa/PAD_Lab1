package socket

import (
	"bufio"
	"io"
	"mr-l0n3lly/go-broker/internal/config"
	"mr-l0n3lly/go-broker/pkg/logging"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Server struct {
	sock net.Listener
}

func (s *Server) Start() {
	cfg := config.GetConfiguration()
	logger := logging.GetLogger()

	var err error

	s.sock, err = net.Listen("tcp4", cfg.SocketServer.Host+":"+strconv.Itoa(cfg.SocketServer.Port))
	logger.Info(cfg.SocketServer.Host + ":" + strconv.Itoa(cfg.SocketServer.Port))
	if err != nil {
		logger.Fatal("socket listen port failed")
	}

	defer s.sock.Close()

	logger.Info("started to accept connections")
	for {
		conn, err := s.sock.Accept()

		if err != nil {
			logger.Fatal(err)
			continue
		}

		go s.handleClients(conn)
	}
}

func (s *Server) handleClients(conn net.Conn) {
	defer conn.Close()

	var (
		buf    = make([]byte, 1024)
		r      = bufio.NewReader(conn)
		logger = logging.GetLogger()
		// w      = bufio.NewReader(conn)
	)

CONNLOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		switch err {
		case io.EOF:
			break CONNLOOP
		case nil:
			logger.Info("received: ", data)
			if isTransportOver(data) {
				break CONNLOOP
			}
		default:
			logger.Fatal("receive data failed: ", err.Error())
			return
		}
	}
}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}

var instance *Server
var once sync.Once

func GetSocketServer() *Server {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("creating new socket server")

		instance = &Server{}
	})

	return instance
}
