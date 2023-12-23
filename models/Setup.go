package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host = localhost user = postgres dbname=train_data port=5432 sslmode=disable  password=root"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to Connect to the database")

	}
	database.AutoMigrate(&Post{})

	DB = database

}
