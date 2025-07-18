package database

import (
	"fmt"
	"log"
	"to-do-list/config"
	"to-do-list/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.Get("DB_HOST"),
		config.Get("DB_USER"),
		config.Get("DB_PASS"),
		config.Get("DB_NAME"),
		config.Get("DB_PORT"),
		config.Get("TIME_ZONE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatal("failed to migrate:", err)
	}
}
