package controller

import (
	"encoding/json"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"github.com/TheNestDevs/SnTx/api/internal/domain"
)

func WebsocketRoute(hub *domain.Hub) fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		var msg domain.Message

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

		hub.NewClient <- &client

		for {
			_, payLoad, _ := conn.ReadMessage()

			_ = json.Unmarshal(payLoad, &msg)

			// sending data to go routine
			hub.Broadcast <- &msg
		}
	})
}
