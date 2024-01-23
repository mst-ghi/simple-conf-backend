package seeder

import (
	"video-conf/core/config"
	"video-conf/database"
)

func Seeder() {
	config.InitializeAndSet()

	database.InitializeAndConnect()
	db := database.Connection()

	UserSeeder(db)
}
