package WebApi

import "github.com/gin-gonic/gin"

func GetAllProjects(c *gin.Context) {
	c.JSON(200, gin.H{
		"get": "allProjects",
	})
}

func ShowPrueba(c *gin.Context) {
	c.JSON(200, gin.H{
		"laucio": "genius",
	})
}
