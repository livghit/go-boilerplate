package routes

import ()

// Defining the Route type with all its params
type Route struct {
	routeType     string
	routePath     string
	routeFunction func()
}

var ROUTES []Route

func LoadRoutes() {
	var routeFunc = func() {
		println("Hello")
	}
	route := Route{routeType: "GET", routePath: "/", routeFunction: routeFunc}
	ROUTES = append(ROUTES, route)
}

func PrintRoutes() {
	println("The Route is : ", ROUTES)
}
