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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
		log.Println(c.Request.RequestURI)
	}
}
