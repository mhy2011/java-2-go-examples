package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/json", json)
	r.Run()	//默认监听8080端口
}

func json(c *gin.Context)  {
	data := map[string]interface{}{
		"lang": "go语言",
		"tag":  "计算机",
	}
	c.AsciiJSON(http.StatusOK, data)

}