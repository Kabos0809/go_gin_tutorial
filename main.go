package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kabos0809/go_gin_tutorial/controller"
	"github.com/kabos0809/go_gin_tutorial/middleware"
)

func main() {
	e := gin.Default()
	e.Use(middleware.RecordUaAndTime)
	//CRUD
	be := engine.Group("/book")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete", controller.BookDelete)
		}
	}
	e.Run(":3000")
}