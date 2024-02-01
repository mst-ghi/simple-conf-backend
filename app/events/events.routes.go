package events

import (
	"video-conf/core/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewEventsController()

	guestGroup := router.Group("/events")
	{
		guestGroup.GET("", ctrl.FindAll)
		guestGroup.GET("/:id", ctrl.FindOne)
	}

	authGroup := router.Group("/events").Use(middlewares.Auth)
	{
		authGroup.POST("", ctrl.Create)
		authGroup.PUT("/:id", ctrl.Update)
	}
}
