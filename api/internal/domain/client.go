package domain

import "github.com/gofiber/contrib/websocket"

// Client holds the structure of a single client instance
type Client struct {
	Id    string          `json:"id"`    // client id fetched from query or if not passed then generated automatically
	Topic string          `json:"topic"` // topic of the conversation
	Conn  *websocket.Conn // websocket connection for each client
}
