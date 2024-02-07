package app

import (
	"video-conf/app/auth"
	"video-conf/app/communities"
	"video-conf/app/events"
	"video-conf/app/messages"
	"video-conf/app/rooms"
	"video-conf/app/users"
	"video-conf/core"

	"github.com/gin-gonic/gin"
)

// @tags App
// @router	/api [get]
// @summary	app route, get heathy status
func HomeRoute(c *gin.Context) {
	ctrl := core.GetController()
	ctrl.Success(c, map[string]string{
		"heathy": "I'm OK...",
	})
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", HomeRoute)

	v1Group := router.Group("/v1")
	{
		auth.RegisterRoutes(v1Group)
		users.RegisterRoutes(v1Group)
		communities.RegisterRoutes(v1Group)
		events.RegisterRoutes(v1Group)
		rooms.RegisterRoutes(v1Group)
		messages.RegisterRoutes(v1Group)
	}
}
