package WebApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/laucio/Entity"
	"github.com/laucio/Repository"
	"log"
	"regexp"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func GetAllProjects(c *gin.Context) {
	projects, error := Repository.GetAllProjects()
	if error != nil {
		log.Println(error)
		c.JSON(500, gin.H{
			"error": fmt.Errorf("%s", error),
		})
	} else {
		c.JSON(200, gin.H{
			"projects": projects,
		})
	}
}

func GetWrongNameProjects(c *gin.Context) {
	pattern := c.Param("pattern")
	projects, error := Repository.GetAllProjects()
	if error != nil {
		log.Println(error)
		c.JSON(500, gin.H{
			"error": fmt.Errorf("%s", error),
		})
	}

	var wrongNameProjects []Entity.Project

	for _, project := range projects {

		match, _ := regexp.MatchString(pattern, project.Name)
		if !match {
			wrongNameProjects = append(wrongNameProjects, project)
		}
	}

	c.JSON(200, gin.H{
		"projects": wrongNameProjects,
	})
}

func GetTimeWindowProjects(c *gin.Context) {
	startDate, _ := time.Parse(layoutISO, c.Param("startdate"))
	endDate, _ := time.Parse(layoutISO, c.Param("enddate"))

	projects, error := Repository.GetAllProjects()
	if error != nil {
		log.Println(error)
		c.JSON(500, gin.H{
			"error": fmt.Errorf("%s", error),
		})
	}

	var timeWindowProjects []Entity.Project

	for _, project := range projects {

		projectStartDate, _ := time.Parse(layoutISO, project.StartDate[0:10])
		if isInTimeWindow(projectStartDate, startDate, endDate) {
			timeWindowProjects = append(timeWindowProjects, project)
		}
	}

	c.JSON(200, gin.H{
		"projects": timeWindowProjects,
	})
}

func GetReadmeProjects(c *gin.Context) {

	projects, error := Repository.GetAllProjects()

	if error != nil {
		c.JSON(500, gin.H{
			"error": fmt.Errorf("%s", error),
		})
	}

	var readmeProjects []Entity.Project

	for _, project := range projects {

		if project.ReadmeUrl != "" {
			readmeProjects = append(readmeProjects, project)
		}
	}

	c.JSON(200, gin.H{
		"projects": readmeProjects,
	})
}

func isInTimeWindow(target, startDate, endDate time.Time) bool {
	if !target.Before(startDate) && target.Before(endDate) {
		return true
	} else {
		return false
	}
}
