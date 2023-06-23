package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-boilerplate/app/database"
	"github.com/livghit/go-boilerplate/app/router"
)

func main() {

	app := fiber.New()
	router.SetupRoutes(app)
  database.CreateConnection()

  app.Listen(":2020")

}
