package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-boilerplate/app/http/controllers"
)

// Router and Routes setup
func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"ping": "pong"})
	})

	app.Get("/users", controllers.GetUsers)

}
