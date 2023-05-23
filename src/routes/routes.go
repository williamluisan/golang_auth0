package routes

import (
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	controller "github.com/williamluisan/golang_auth0/controllers"
	middlewares "github.com/williamluisan/golang_auth0/middlewares"
)

func Routes(router *gin.Engine) {
	router.GET("/token", controller.GetToken)
	
	userRoute := router.Group("user") 
	{
		userRoute.GET("/list", controller.GetListUser)
	}
	
	activitiesRoute := router.Group("activitites") 
	{
		userRoute.Use(adapter.Wrap(middlewares.EnsureValidToken()))
		activitiesRoute.GET("/list", controller.GetListActivities)
	}
}