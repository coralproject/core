package main

import (
	"log"

	"github.com/coralproject/core/app/models"
	"github.com/coralproject/core/database/postgres"
	"github.com/jinzhu/gorm"
)

func migrate(db *gorm.DB) {

	db.LogMode(true)

	db.AutoMigrate(&models.Article{})

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Recommendation{})
}

func main() {

	//Connection to the DB
	db, err := postgres.Open()

	if err != nil {
		log.Fatal("Could not connect to the database.")
	}
	defer postgres.Close(db)

	migrate(db)
}
