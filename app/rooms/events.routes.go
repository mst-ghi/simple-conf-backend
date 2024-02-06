package rooms

import (
	"video-conf/core/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewRoomsController()

	authGroup := router.Group("/rooms").Use(middlewares.Auth)
	{
		authGroup.GET("", ctrl.FindAll)
		authGroup.GET("/:id", ctrl.FindOne)
		authGroup.POST("", ctrl.Create)
		authGroup.PUT("/:id", ctrl.Update)
	}
}
