package router

import "github.com/gin-gonic/gin"

type Router struct {
	routes []Route
	engine *gin.Engine
}


func NewRouter() *Router {
	return NewRouter().initRouter()
}

func (r *Router) LoadRoutes() {
	r.routes = Routes
}

func (r *Router) initRouter() *Router {
	r.engine = gin.Default()

	for _, route := range r.routes {
		rt := route.routeType
		rp := route.routeParam
		rpath := route.routePath
		if rt == "GET" {
			r.engine.GET(rpath, rp)
		} else {
			r.engine.POST(rpath, rp)
		}
	}

	return r
}
