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

		defer func() {
			hub.RemoveClient(&client)
			_ = conn.Close()
		}()

		for {
			messageType, payLoad, _ := conn.ReadMessage()

			_ = json.Unmarshal(payLoad, &msg)

			// broadcasting message to all other clients in the same room
			if messageType == 1 { // checking if the message is a text type message
				hub.Broadcast <- &msg
			} else {
				break
			}
		}
	})
}
