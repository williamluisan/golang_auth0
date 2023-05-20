package controllers

import "github.com/gin-gonic/gin"

func GetListUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list of users",
	})
}