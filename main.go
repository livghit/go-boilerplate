package main

import "github.com/livghit/go-boilerplate/app/router"
import "github.com/gin-gonic/gin"

func main() {
	r := router.NewRouter()

	r.AddRoute("GET", "/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.AddRoute("GET", "/pong", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.RouterServe()
}
