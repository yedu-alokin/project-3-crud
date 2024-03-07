package utils

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, httpCode int, errors interface{}) {
	go Looger(c, errors)

	c.JSON(httpCode, gin.H{
		"message": "An error happened",
		"error":   errors,
	})
}

func SuccessResponse(c *gin.Context, httpCode int, data interface{}) {
	go Looger(c, "success")

	c.JSON(httpCode, gin.H{
		"message": "Success",
		"data":    data,
	})
}
