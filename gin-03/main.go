package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 定义新路由，返回json类型数据
	router.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"result":    "ok",
			"data":      "Hello kugou",
			"developer": "bee",
		})
	})

	router.Run(":8090")
}
