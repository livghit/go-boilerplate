package router

import "github.com/gin-gonic/gin"

type Route struct {
	routeType  string
	routePath  string
	routeParam func(*gin.Context)
}

var Routes []Route

func newRoute(rt string, rpath string, rp func(*gin.Context)) *Route {
	return &Route{routeType: rt, routePath: rpath, routeParam: rp}
}

func (r *Router) AddRoute(routeType string, rpath string, routeParam func(*gin.Context)) {
	rt := routeType
	rp := routeParam
	route := newRoute(rt, rpath, rp)
	Routes = append(Routes, *route)
}

func (r *Route) Routes() []Route {
	return Routes
}
