package router 

import "github.com/gin-gonic/gin"

type Router struct{
  routes []Route
  engine gin.Engine 
}


func (r *Router) LoadRoutes(){
  r.routes = Routes
}


