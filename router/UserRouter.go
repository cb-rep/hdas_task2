package router

import (
	"github.com/gin-gonic/gin"
	"hdas_task/controller"
)

func Start() {
	e := gin.Default()

	// 导入templates下的html文件
	e.LoadHTMLGlob("templates/*")

	//导入assets包下的文件
	e.Static("/assets","./assets")

	e.POST("/register",controller.Register)
	e.GET("/register",controller.GoRegister)

	e.GET("/login",controller.GoLogin)
	e.POST("/login",controller.Login)

	e.GET("/",controller.Index)

	e.GET("/books",controller.GetAllBooks)

	e.POST("/addBook",controller.AddBook)
	e.GET("/addBook",controller.GoAddBook)

	e.GET("/user",controller.GetUser)
	e.POST("/user",controller.GetUser)

	e.Run()
}