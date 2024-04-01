package bootstrap

import (
	"simple-conf/app/gateway"
	"simple-conf/core"
	"simple-conf/core/config"
	"simple-conf/core/engine"
	"simple-conf/core/scheduler"
	"simple-conf/database"
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
