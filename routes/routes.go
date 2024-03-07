package routes

import "github.com/gin-gonic/gin"

func StartRouting(g *gin.Engine) {

	Students(g.Group("/students"))
}
