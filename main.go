package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 加载静态文件
	r.Static("/xxx", "/Users/luliang/GoLand/homePage/assets")
	r.LoadHTMLGlob("/Users/luliang/GoLand/homePage/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":9092")
}
