package app

import (
	"simple-conf/app/auth"
	"simple-conf/app/comments"
	"simple-conf/app/communities"
	"simple-conf/app/events"
	"simple-conf/app/messages"
	"simple-conf/app/rooms"
	"simple-conf/app/users"
	"simple-conf/core"

	"github.com/gin-gonic/gin"
)

// @tags    App
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
		comments.RegisterRoutes(v1Group)
	}
}
