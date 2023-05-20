package routes

import (
	"github.com/gin-gonic/gin"
	CActivity "github.com/williamluisan/golang_auth0/controllers/activities"
	CUser "github.com/williamluisan/golang_auth0/controllers/users"
)

func Routes(router *gin.Engine) {
	router.GET("/users", CUser.GetListUser)
	router.GET("/activity", CActivity.GetListActivities)
}