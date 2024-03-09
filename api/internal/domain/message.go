package domain

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

var (
	Cli     = make(chan Client) // Cli is client channel for communicating with go routine
	PayLoad = make(chan []byte) // PayLoad channel contains the payload sent from client
	Svr     = make(chan []*Client)
	ctx     = context.Background() // ctx is used as context passed to redis
)

// Client holds the structure of a single client instance
type Client struct {
	Id    string          `json:"id"`    // client id fetched from query or if not passed then generated automatically
	Topic string          `json:"topic"` // topic of the conversation
	Conn  *websocket.Conn // websocket connection for each client
}

type Server struct {
	server []*Client
}

// Message holds the structure of JSON message send via websocket.
type Message struct {
	ClientID string `json:"client_id,omitempty"` // client id of the client
	Topic    string `json:"topic,omitempty"`     // topic of the message sent
	Msg      string `json:"message,omitempty"`   // message string that is sent
}

func (s *Server) Send(client *Client, msg Message) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	err = client.Conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Fatalf("websocket error: %v", err)
		}
	}
}

func (s *Server) ProcessMessage() {
	for {
		select {
		case client := <-Cli:
			s.LoadServer(&client)
		case payload := <-PayLoad:
			var msg Message
			err := json.Unmarshal(payload, &msg)
			if err != nil {
				log.Fatalf("error unmarshalling payload: %v", err)
			}

			for _, client := range s.server {
				if client.Id == msg.ClientID {
					continue
				} else if client.Topic == msg.Topic {
					s.Send(client, msg)
				}
			}
		}
	}
}

func (s *Server) LoadServer(client *Client) {
	s.server = append(s.server, client)
}
