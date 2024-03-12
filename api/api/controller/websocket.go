package controller

import (
	"encoding/json"
	"fmt"

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

			_, payLoad, _ := conn.ReadMessage()

			_ = json.Unmarshal(payLoad, &msg)

			fmt.Println(msg)

			// broadcasting message to all other clients in the same room
			hub.Broadcast <- &msg
		}

	})
}
