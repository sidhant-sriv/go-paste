package db

import (
	"sidhant/pastebin-clone/models"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB connection
func GetDB() (*gorm.DB, error) {
	//setup db
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=pastebin password=admin sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// init db and returns db and error. User and Pastes models are migrated
func InitDB() (*gorm.DB, error) {
	//setup db
	db, err := GetDB()
	if err != nil {
		return nil, err
	}
	//migrate models
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Paste{})
	return db, nil
}
