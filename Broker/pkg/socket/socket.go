package socket

import (
	"bufio"
	"encoding/json"
	"io"
	"mr-l0n3lly/go-broker/internal/client"
	"mr-l0n3lly/go-broker/internal/config"
	"mr-l0n3lly/go-broker/internal/db"
	"mr-l0n3lly/go-broker/internal/messages"
	"mr-l0n3lly/go-broker/pkg/logging"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Server struct {
	sock net.Listener
	DB   db.Database
}

func (s *Server) Start(db db.Database) {
	s.DB = db
	cfg := config.GetConfiguration()
	logger := logging.GetLogger()

	var err error

	addr := net.TCPAddr{
		Port: cfg.SocketServer.Port,
		IP:   net.ParseIP(cfg.SocketServer.Host),
	}

	s.sock, err = net.ListenTCP("tcp", &addr)
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
	var (
		buf    = make([]byte, 1024)
		r      = bufio.NewReader(conn)
		logger = logging.GetLogger()
		data   = make([]byte, 1)
		//w      = bufio.NewWriter(conn)
	)

CONNLOOP:
	for {
		n, err := r.Read(buf)
		data = append(data, buf[:n]...)

		switch err {
		case io.EOF:
			logger.Error("%v", err)
			defer conn.Close()
			break CONNLOOP
		case nil:
			logger.Info("received: ", string(data))
			if isTransportOver(string(data)) {
				message, leftOver, _ := strings.Cut(string(data), "\r\n\r\n")
				jsonData := messages.SenderRequest{}

				if message[0] == 0 {
					err = json.Unmarshal([]byte(message)[1:], &jsonData)
				} else {
					err = json.Unmarshal([]byte(message)[0:], &jsonData)
				}

				data = []byte(leftOver)

				if err != nil {
					logger.Error("%v", err)
					defer conn.Close()
					break CONNLOOP
				}

				messagesHandler := client.Handler{
					DB: s.DB,
				}

				response, err := messagesHandler.ParseMessage(jsonData, &conn)

				if err != nil {
					logger.Error("%v", err)
				}

				logger.Info(string(response))

				_, err = conn.Write(response)

				if err != nil {
					logger.Error("%v", err)
				}

				//if jsonData.Action != messages.SubscribeAction {
				//	defer conn.Close()
				//}

				// break CONNLOOP
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
