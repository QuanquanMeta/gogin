package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Test

func main() {
	r := gin.Default()

	// get
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%v", "Front page")
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
			"msg":     "hello gin",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "hello gin",
		})
	})
	// post
	r.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%v", "POST")
	})
	// put
	r.PUT("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%v", "PUT Edit")
	})
	// delete
	r.DELETE("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "value:%v", "DELETE")
	})

	r.Run(":8800") // start a web service
}
