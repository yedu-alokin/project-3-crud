package routes

import (
	"crud/handlers"

	"github.com/gin-gonic/gin"
)

func Students(router *gin.RouterGroup) {
	router.GET("/", handlers.GetAllStudents)
	router.GET("/:id", handlers.GetStudent)
	router.POST("/", handlers.AddNewStudent)
	router.PUT("/:id", handlers.UpdateStudent)
	router.DELETE("/:id", handlers.DeleteStudent)

}
