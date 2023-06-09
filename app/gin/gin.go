package main

import "github.com/gin-gonic/gin"
import "github.com/livghit/go-boilerplate/app/routes"

type Router struct{
  routes []routes.Route
  engine gin.Engine 
}


func (r *Router) LoadRoutes(){
  r.routes = routes.Routes
}
