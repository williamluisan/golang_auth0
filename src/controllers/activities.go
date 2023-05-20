package controllers

import "github.com/gin-gonic/gin"

func GetListActivities(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list of activities",
	})
}