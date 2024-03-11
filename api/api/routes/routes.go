package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/TheNestDevs/SnTx/api/api/controller"
	"github.com/TheNestDevs/SnTx/api/internal/domain"
)

func SetupRoute(app *fiber.App) {
	ws := app.Group("/ws")

	hub := domain.NewHub()

	go hub.Run()
	ws.Get("/:id", controller.WebsocketRoute(hub))
}
