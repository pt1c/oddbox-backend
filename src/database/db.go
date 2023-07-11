package database

import (
	"fmt"
	"log"
	"oddbox/src/config"
	"oddbox/src/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func Connect() {
	var err error
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Database successfully connected.")
}

func AutoMigrate() {
	DB.AutoMigrate(
		models.User{},
		models.Item{},
		models.Item_Type{},
	)
}