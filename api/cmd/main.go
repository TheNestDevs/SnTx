package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/TheNestDevs/SnTx/api/internal/config"
)

func main() {
	appConfig := config.NewAppConfig()

	app := fiber.New()

	app.Use(cors.New())

	if err := app.Listen(":" + appConfig.AppConfig.Port); err != nil {
		panic(err)
	}
}
