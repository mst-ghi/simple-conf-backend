package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewUsersController()

	guestGroup := router.Group("/users")
	{
		guestGroup.GET("", ctrl.FindAll)
		guestGroup.GET("/:id", ctrl.FindOne)
	}
}
