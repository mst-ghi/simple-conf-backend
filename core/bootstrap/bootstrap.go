package bootstrap

import (
	"video-conf/app/gateway"
	"video-conf/core"
	"video-conf/core/config"
	"video-conf/core/engine"
	"video-conf/core/scheduler"
	"video-conf/database"
)

func Serve() {
	config.InitializeAndSet()
	database.InitializeAndConnect()

	core.Initialize()
	engine.Initialize()

	gateway.Initialize()
	gateway.Serve(engine.GetEngine())

	engine.RegisterRoutes()

	go scheduler.InitScheduler()

	defer gateway.GetSocket().Close()
	engine.Serve()
}
