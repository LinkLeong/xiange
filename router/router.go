package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	gin.SetMode("debug")
	//设置静态文件路径
	//r.Static("/assets/", "./assets")
	r.GET("/text/:c/index.html", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	//r.GET("/login", view.Login)
	//r.POST("/login", view.LoginP)

	//r.GET("/welcome", view.Welcome)
	//
	//memberGroup := r.Group("/member")
	//memberGroup.Use()
	//{
	//	memberGroup.GET("/list", member.List)
	//}

	return r

}

