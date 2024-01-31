package engine

import (
	"video-conf/app"
	"video-conf/core/config"
	"video-conf/core/middlewares"
	"video-conf/core/swagger"
	"video-conf/pkg/handlers"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func GetEngine() *gin.Engine {
	return engine
}

func Initialize() {
	gin.SetMode(config.Get("GIN_MODE"))

	engine = gin.Default()
	engine.SetTrustedProxies(nil)
	engine.RedirectTrailingSlash = true
	engine.RedirectFixedPath = true

	engine.Use(gin.CustomRecovery(handlers.InternalErrorHandler))
	engine.Use(middlewares.Cors())
}

func Serve(addr ...string) {
	runAddress := config.GetRunAddress()

	if addr != nil {
		runAddress = addr[0]
	}

	engine.Run(runAddress)
}

func RegisterRoutes() {
	routerGroup := GetEngine().Group("api")

	app.RegisterRoutes(routerGroup)

	swagger.RegisterSwagger(routerGroup)
}
