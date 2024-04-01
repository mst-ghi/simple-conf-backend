package seeder

import (
	"simple-conf/core/config"
	"simple-conf/database"
)

func Seeder() {
	config.InitializeAndSet()

	database.InitializeAndConnect()
	db := database.Connection()

	UserSeeder(db)
}
