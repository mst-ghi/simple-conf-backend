package comments

import (
	"video-conf/core/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewCommentsController()

	guestGroup := router.Group("/events")
	{
		guestGroup.GET("", ctrl.FindAll)
	}

	authGroup := router.Group("/events").Use(middlewares.Auth)
	{
		authGroup.POST("", ctrl.Create)
		authGroup.PUT("/:id", ctrl.Update)
	}
}
