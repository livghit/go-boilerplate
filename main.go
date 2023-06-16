package main

import (
	"github.com/livghit/go-boilerplate/routes"
)

func main() {
	// Main Func to run the server

	routes.PrintRoutes()
	routes.LoadRoutes()
	routes.PrintRoutes()
}
