package bootstrap

import (
	"video-conf/core"
	"video-conf/core/config"
	"video-conf/core/engine"
	"video-conf/database"
)

func Serve() {
	config.InitializeAndSet()
	database.InitializeAndConnect()

	core.Initialize()

	engine.Initialize()
	engine.RegisterRoutes()
	engine.Serve()
}
