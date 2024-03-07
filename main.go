package main

import (
	"crud/routes"
	"crud/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server running...",
		})
	})
	utils.ConnectDatabase()
	defer utils.DbClient.Close()

	routes.StartRouting(r)

	
	r.Run() // listen and serve on 0.0.0.0:8080
}
