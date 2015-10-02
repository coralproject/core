package postgres

import (
	"log"

	"github.com/coralproject/core/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Open() (*gorm.DB, error) {
	// get the config
	c := config.GetDatabase()

	// build the connect string for postgres
	s := getPostgresqlConnection(c)

	// open, or so it appears, but not yet
	db, err := gorm.Open(c.Adapter, s)
	if err != nil {
		log.Fatal("Could not connect to PostgreSQL with ", s, err)
		return nil, err
	}

	// Trigger actual open via ping. Validate DSN data:
	err = db.DB().Ping()
	if err != nil {
		log.Fatal("Could not connect to the database with", s, err)
		return nil, err
	}

	// set to exported package global
	//Db = &db

	return &db, nil
}

// Close the db connection
func Close(db *gorm.DB) { //error {
	db.Close() // TODO : needs to return an apropiate error
	//return nil
}

func getPostgresqlConnection(c config.DatabaseConfig) string {
	return "user=" + c.Username + " dbname=" + c.Database + " sslmode=disable"
}
