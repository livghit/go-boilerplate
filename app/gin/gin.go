package main

import "github.com/gin-gonic/gin"
import "github.com/livghit/app/routes"

type Router struct{
  routes []routes
  engine gin.Engine 
}


func (r *Router) LoadRoutes(){
  r.engine = *gin.Default()
  routes := routes 
}
