package migrations

import (
	"fmt"
	"log"
	"simple-conf/core/config"
	"simple-conf/database"
	"simple-conf/database/models"
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
		&models.Room{},
		&models.RoomUser{},
		&models.Message{},
		&models.Comment{},
	)

	if err != nil {
		log.Fatal("Cannot migrate")
		return
	}

	fmt.Println("Migration done ..")
}
