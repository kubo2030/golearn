package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	// 定义路由 /form，渲染表单页面
	router.GET("/form", func(ctx *gin.Context) {
		ctx.HTML(200, "form.html", gin.H{})
	})

	// 定义路由 /service,接收POST数据，返回JSON数据类型
	router.POST("/service", func(ctx *gin.Context) {
		uname := ctx.PostForm("uname")
		ctx.JSON(200, gin.H{
			"result": "ok",
			"hello":  uname,
		})
	})

	router.Run(":8090")
}
