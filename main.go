package main

import (
	"github.com/firdausalif/go-fiber/database"
	"github.com/firdausalif/go-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	router.SetupRoutes(app)

	app.Listen(":3000")
}
