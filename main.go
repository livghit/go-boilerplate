package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-boilerplate/app/http/router"
)

func main() {

	app := fiber.New()
	router.SetupRoutes(app)

  app.Listen(":2020")

}
