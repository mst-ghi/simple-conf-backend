package messages

import (
	"video-conf/core/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewMessagesController()

	authGroup := router.Group("/messages").Use(middlewares.Auth)
	{
		authGroup.GET("", ctrl.FindAll)
		authGroup.POST("", ctrl.Create)
		authGroup.PUT("/:id", ctrl.Update)
	}
}
