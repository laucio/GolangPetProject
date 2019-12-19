package WebApi

import (
	"github.com/gin-gonic/gin"
	"github.com/laucio/Repository"
)

func GetAllProjects(c *gin.Context) {
	data, error := Repository.GetAllProjects()

	if error != nil {
		c.JSON(500, gin.H{
			"error": "ocurrio el error",
		})
	} else {
		c.JSON(200, gin.H{
			"parcialError": 0,
			"projects":     data,
		})
	}
}

func ShowPrueba(c *gin.Context) {
	c.JSON(200, gin.H{
		"laucio": "genius",
	})
}
