package main

import (
	"github.com/gin-gonic/gin"
	"github.com/laucio/WebApi"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	//Public group
	public := r.Group("/public")
	public.GET("/publictest", WebApi.GetAllProjects)

	//Private group
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"laucio": "lauciolaucio2",
	}))
	authorized.GET("/privatetest", WebApi.GetAllProjects)
	authorized.GET("/getallprojects", WebApi.GetAllProjects)
	authorized.GET("/getWrongNameProjects/:pattern", WebApi.GetWrongNameProjects)
	authorized.GET("/getTimeWindowProjects/:startdate/:enddate", WebApi.GetTimeWindowProjects)
	authorized.GET("/getReadmeProjects", WebApi.GetReadmeProjects)
	r.Run()
}

// docker build -t petprojectfinalversion .
// Docker docker run -d -p 5000:8080 petprojectfinalversion
// Corre en localhost:5000
