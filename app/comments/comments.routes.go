package comments

import (
	"simple-conf/core/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewCommentsController()

	guestGroup := router.Group("/comments")
	{
		guestGroup.GET("", ctrl.FindAll)
	}

	authGroup := router.Group("/comments").Use(middlewares.Auth)
	{
		authGroup.POST("", ctrl.Create)
		authGroup.PUT("/:id", ctrl.Update)
	}
}
