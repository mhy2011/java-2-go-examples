package main

import (
	"github.com/gin-gonic/gin"
	"java-2-go-examples/gin/controller"
	"java-2-go-examples/gin/req"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/json", json)
	r.GET("/index", index)
	r.POST("/login", login)
	r.POST("/loginJson", controller.LoginJson)
	r.Run()	//默认监听8080端口
}

func json(c *gin.Context)  {
	data := map[string]interface{}{
		"lang": "go语言",
		"tag":  "计算机",
	}
	c.AsciiJSON(http.StatusOK, data)

}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "gin测试",
	})
}

// 登录
func login(c *gin.Context) {
	var form req.LoginForm
	if c.ShouldBind(&form) == nil {
		if form.Username == "root" && form.Password == "root" {
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "invalid username or password",
			})
		}
	}
}
