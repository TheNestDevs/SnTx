package main

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/TheNestDevs/SnTx/api/api/routes"
	"github.com/TheNestDevs/SnTx/api/internal/config"
)

func main() {
	appConfig := config.NewAppConfig()

	app := fiber.New()

	app.Use(cors.New())
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	routes.SetupRoute(app)

	if err := app.Listen(":" + appConfig.AppConfig.Port); err != nil {
		panic(err)
	}
}
