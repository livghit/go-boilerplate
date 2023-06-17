package router

import (
	"github.com/gofiber/fiber/v2"
  "github.com/livghit/go-boilerplate/app/http/controllers"
)

// Router and Routes setup
func SetupRoutes(app *fiber.App) {

	// Hello World route

	app.Get("/", controllers.GetUsers)

}
