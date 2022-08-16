package main

import (
	"github.com/gin-gonic/gin"
	// "log"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"test-golang/controller"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", RootHandler)
	api.POST("/tis", PostHandler)
	api.GET("/test/:id", RootHandler)
	router.Run(":8888")
}
