package mysql

import (
	"log"

	"github.com/coralproject/core/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Open() (*gorm.DB, error) {
	// get the config
	c := config.GetDatabase()

	// build the connect string for postgres
	s := getMySQLConnection(c)

	// open, or so it appears, but not yet
	db, err := gorm.Open(c.Adapter, s)
	if err != nil {
		log.Fatal("Could not connect to MySQL with ", s, err)
		return nil, err
	}

	// Trigger actual open via ping. Validate DSN data:
	err = db.DB().Ping()
	if err != nil {
		log.Fatal("Could not connect to the database with", s, err)
		return nil, err
	}

	// set to exported package global
	//Db = db

	return &db, nil
}

// Close the db connection
func Close(db *gorm.DB) { //error {
	db.Close() // TODO : needs to return an apropiate error
	//return nil
}

func getMySQLConnection(c config.DatabaseConfig) string {
	return c.Username + ":" + c.Password + "@" + c.Host + "/" + c.Database
}
