package database

import (
	"fmt"
	"log"
	"user-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() (db *gorm.DB, err error) {
	config := config.GetDBConfig()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=%s",
		config.Username, config.Password, config.Host, config.Port, config.Name, config.TimeZone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}
