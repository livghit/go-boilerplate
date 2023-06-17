package router

import (
	"github.com/gofiber/fiber/v2"
)

// Router and Routes setup
func SetupRoutes(app *fiber.App) {

	// Hello World route

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

}
