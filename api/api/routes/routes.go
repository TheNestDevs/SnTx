package routes

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"github.com/TheNestDevs/SnTx/api/api/controller"
)

func SetupRoute(app *fiber.App) {
	ws := app.Group("/ws")
	go controller.Server.ProcessMessage()
	ws.Get("/:id", websocket.New(controller.WebsocketRoute))
}
