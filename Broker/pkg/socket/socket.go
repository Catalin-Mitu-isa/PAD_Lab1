package socket

import (
	"bufio"
	"encoding/json"
	"io"
	"mr-l0n3lly/go-broker/internal/config"
	"mr-l0n3lly/go-broker/internal/db"
	"mr-l0n3lly/go-broker/internal/messages"
	"mr-l0n3lly/go-broker/internal/models"
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
		w      = bufio.NewWriter(conn)
	)

CONNLOOP:
	for {
		n, err := r.Read(buf)
		data := buf[:n]
		json_data := messages.SenderMessage{}
		err = json.Unmarshal(data, &json_data)

		switch json_data.Action {

		case messages.CREATE_TOPIC_ACTION:
			correct_data := messages.SenderCreateRequest{}
			json.Unmarshal(data, &correct_data)

			// Add topic to database
			s.DB.AddTopic(models.Topic{
				TopicName: correct_data.TopicName,
			})

			// Craft a response for sender
			response := messages.SenderCreateResponse{
				SenderResponse: messages.SenderResponse{
					SenderMessage: messages.SenderMessage{
						Action: messages.CREATE_TOPIC_ACTION,
					},
					Success: true,
					Error:   "",
				},
			}

			response_json, _ := json.Marshal(response)
			w.Write(response_json)

		case messages.PUBLISH_MESSAGE_ACTION:
			correct_data := messages.SenderCreateRequest{}
			json.Unmarshal(data, &correct_data)
		}

		switch err {
		case io.EOF:
			break CONNLOOP
		case nil:
			logger.Info("received: ", data)
			if isTransportOver(string(data)) {
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
