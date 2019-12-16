package main

import (
	"github.com/gin-gonic/gin"
	"github.com/laucio/WebApi"
	"io"
	"os"
)

func main() {
	r := gin.Default()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//Public group
	public := r.Group("/public")
	public.GET("/publictest", WebApi.ShowPrueba)

	//Private group
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"laucio": "lauciolaucio2",
	}))
	authorized.GET("/privatetest", WebApi.ShowPrueba)
	authorized.GET("/getallprojects", WebApi.GetAllProjects)
	r.Run()
}
