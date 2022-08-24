package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Test

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%v", "hello go gin")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%v", "POST")
	})

	r.PUT("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%v", "PUT Edit")
	})

	r.DELETE("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%v", "DELETE")
	})

	r.Run(":8800") // start a web service
}
