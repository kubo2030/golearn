package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/data", func(ctx *gin.Context) {
		ctx.HTML(200, "data.html", gin.H{"data": "Hello Go/Gin World"})
	})
	router.Run(":8090")
}
