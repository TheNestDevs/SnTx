package controller

import (
	"github.com/gofiber/contrib/websocket"

	"github.com/TheNestDevs/SnTx/api/internal/domain"
)

var (
	Server domain.Server
)

func WebsocketRoute(conn *websocket.Conn) {
	id := conn.Params("id")
	if id == "" {
		_ = conn.Close()
	}

	topic := conn.Query("topic")
	if topic == "" {
		_ = conn.Close()
	}

	client := domain.Client{
		Id:    id,
		Topic: topic,
		Conn:  conn,
	}

	for {
		_, payLoad, _ := conn.ReadMessage()

		// sending data to go routine
		domain.Cli <- client
		domain.PayLoad <- payLoad
	}
}
