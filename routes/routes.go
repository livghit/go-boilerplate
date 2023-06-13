package routes

import (
	"github.com/gofiber/fiber/v2"
)

// Defining the Route type with all its params
type Route struct {
	routeType     string
	routePath     string
	routeFunction func()
}

func LoadRoutes(router *fiber.App) {

}
