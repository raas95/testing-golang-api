package main

import (
	"github.com/gin-gonic/gin"
	// "log"
	// "github.com/raas95/testing-golang-api/controller"
)

func main() {
	router := gin.Default()
	// api := router.Group("/api")
	// api.GET("/", RootHandler)
	// api.POST("/tis", PostHandler)
	// api.GET("/test/:id", RootHandler)

	router.Run(":8888")
}
