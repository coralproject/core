package main

import (
	"log"

	"github.com/coralproject/core/app/models"
	"github.com/coralproject/core/db/postgres"
	"github.com/gabelula/gorm"
)

// Migrates to the last changes to the models
func migrateAll(db *gorm.DB) {

	//db.LogMode(true)

	db.AutoMigrate(&models.Article{})

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Recommendation{})
}

func migrateByModel(db *gorm.DB, model gorm.Model) {}

func main() {

	// flags
	// models
	// silent
	//flag.Usage = usage // To Do: How to use this flag
	// flag.Parse()
	// args := flag.Args()
	//
	// fmt.Println(args)
	//
	// model_to_migrate := args[0]

	// if model_to_migrate has anything
	// 	migrate all of them
	// 	return

	//migrate all

	//Connection to the DB
	db, err := postgres.Open()

	if err != nil {
		log.Fatal("Could not connect to the database.")
	}
	defer postgres.Close(db)

	migrateAll(db)
}
