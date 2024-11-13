package main

import (
	"github.com/gin-gonic/gin"
	"go-test-basic/common"
	"go-test-basic/handler"
)

func main() {
	common.InitDB()

	r := gin.Default()

	r.POST("/create", handler.CreateItem)
	r.GET("/get", handler.GetItem)
	r.GET("/list", handler.ListItems)
	r.PATCH("/update", handler.UpdateItem)
	r.DELETE("/delete", handler.DeleteItem)

	r.Run()

}
