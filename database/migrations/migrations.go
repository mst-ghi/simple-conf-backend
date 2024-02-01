package migrations

import (
	"fmt"
	"log"
	"video-conf/core/config"
	"video-conf/database"
	"video-conf/database/models"
)

func Migrate() {
	config.InitializeAndSet()

	database.InitializeAndConnect()
	db := database.Connection()

	err := db.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Community{},
		&models.Event{},
	)

	if err != nil {
		log.Fatal("Cannot migrate")
		return
	}

	fmt.Println("Migration done ..")
}
