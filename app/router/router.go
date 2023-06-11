package router

import "github.com/gin-gonic/gin"

type Router struct {
	routes []Route
	engine *gin.Engine
}


func NewRouter() *Router {
	return &Router{routes: Routes , engine: gin.Default()}
}

func InitRouter(r *Router) *Router {

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

func (r *Router) RouterServe(){
  r.engine.Run()
}
