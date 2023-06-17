package main

import (
	"gin/handler"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)
var Router * gin.Engine
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.Use(middleware.ValidateAPIKey())
	r.POST("/login", handler.Login)

	r.Run()
}
