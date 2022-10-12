package models

import (
	"fmt"
	"net"
)

type Subscriber struct {
	TopicName string `json:"topic_name"`
	Conn      *net.Conn
	Addr      net.Addr
}

var subscribers = make([]Subscriber, 0)

func (s Subscriber) AddSubscriber(topicName string, conn *net.Conn) {
	sub := s.GetTopicSubscribers(topicName)

	fmt.Println("Adding subscriber", topicName)
	for _, el := range subscribers {
		if el.Conn == conn {
			return
		}
	}

	subscribers = append(sub, Subscriber{
		TopicName: topicName,
		Conn:      conn,
	})
}

func (s Subscriber) GetAllSubscribers() []Subscriber {
	return subscribers
}

func (s Subscriber) GetTopicSubscribers(topicName string) []Subscriber {
	var topicSubscribers = make([]Subscriber, 0)

	for _, el := range subscribers {
		if el.TopicName == topicName {
			topicSubscribers = append(topicSubscribers, el)
		}
	}

	return topicSubscribers
}

func (s Subscriber) RemoveTopicSubscriber(addr net.Addr) {
	index := -1

	for i, el := range subscribers {
		if el.Addr == addr {
			index = i
		}
	}

	subscribers = append(subscribers[:index], subscribers[index+1])
}
