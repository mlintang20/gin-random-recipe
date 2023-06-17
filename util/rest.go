package util

import "github.com/gin-gonic/gin"

func BuildResponse(message string, data interface{}) gin.H {
	return gin.H{
		"message": message,
		"data":    data,
	}
}