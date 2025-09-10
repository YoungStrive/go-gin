package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	Account string `form:"account"`
	Pwd     string `form:"pwd"`
	//多个逗号分隔
	Email string `json:"email" binding:"required,email"`
}

// 中间件1
func mw1() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mw1 before")
		c.Next()
		fmt.Println("mw1 after")
	}
}
func mw2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("mw2 before")
		c.Next()
		fmt.Println("mw2 after")
	}
}
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

	//获取静态文件
	router.Static("/static", "./static")
	//不显示文件名
	router.StaticFile("/f1", "./static/1.text")

	//需要先加载模板
	router.LoadHTMLGlob("templates/**/*")

	router.GET("/index1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "top/index.impl", gin.H{"title": "我是首页！"})
	})

	//占位的
	router.GET("/index/:u/:a", func(c *gin.Context) {
		name := c.Param("u")
		age := c.Param("a")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})

	})
	router.GET("/index?name=u1&age=a1", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})

	})
	//post  form-data
	router.POST("/login", func(c *gin.Context) {
		a := c.PostForm("account")
		p := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"account": a,
			"pwd":     p,
		})
	})

	//post  json 结构体获取
	router.POST("/loginJson", func(c *gin.Context) {
		info := &LoginInfo{}
		err := c.ShouldBind(info)
		if err != nil {
			c.JSON(200, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, info)
	})

	//测试中间件
	router.GET("/testNw", mw1(), mw2(), func(c *gin.Context) {
		fmt.Println("我是接口的")
		c.String(200, "Hello MY IS zj")

	})
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"yl": "123456",
	}))
	//再次测试 当loclhost:8090/admin/testLogin  就要认证
	authorized.GET("/testLogin", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		c.String(200, user)

	})

	port := ":8090"
	router.Run(port)
}
