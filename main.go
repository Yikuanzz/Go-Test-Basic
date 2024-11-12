package main

import (
	"go-test-basic/common"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	r := gin.Default()

	r.Run()

}
