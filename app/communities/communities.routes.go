package communities

import (
	"video-conf/core/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewCommunitiesController()

	guestGroup := router.Group("/communities")
	{
		guestGroup.GET("", ctrl.FindAll)
		guestGroup.GET("/:id", ctrl.FindOne)
	}

	authGroup := router.Group("/communities").Use(middlewares.Auth)
	{
		authGroup.POST("/", ctrl.Create)
		authGroup.PUT("/:id", ctrl.Update)
		authGroup.PUT("/:id/join", ctrl.Join)
	}
}
