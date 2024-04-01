package database

import (
	"log"
	"os"
	"simple-conf/core/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connection() *gorm.DB {
	return DB
}

func InitializeAndConnect() {
	dialector := mysql.Open(config.GetDatabaseDNS())

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  true,
			},
		),
	})

	if err != nil {
		panic("Cannot connect to database")
	}

	DB = db
}
