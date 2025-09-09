package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//路由
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Winnie",
		})
	})

	//路由分组
	group1 := router.Group("/v1")
	{
		group1.GET("/user", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "我是V1",
			})
		})
	}
	//路由分组
	group2 := router.Group("/v2")
	{
		group2.GET("/user", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "我是V2",
			})
		})
	}

	//重定向
	router.GET("test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")

	})
	// 重定向到内部
	router.GET("/test1", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/ping")
	})
	port := ":8090"
	router.Run(port)
}
