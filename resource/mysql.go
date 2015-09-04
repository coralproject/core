package mysql

/*
Package resource/mysql handles the mysql db connection

Todo:
* Add support for master/slave connections
* Add support for multi environments
*/

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/coralproject/core/config"
	"github.com/coralproject/core/log"
)

// Db points to the mysql db
var Db *sql.DB

// Open the db connection using coralproject/core/config creds
func Open() *sql.DB {

	// get the config
	c := config.GetMySQL()

	// build the connect string for MySQL
	s := connectString(c)

	// open, or so it appears, but not yet
	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal("Could not connect to MySQL with", s, err)
	}

	// Trigger actual open via ping. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to MySQL with", s, err)
	}

	// set to exported package global
	Db = db

	return Db

}

// Close the db connection
func Close() {
	Db.Close()
}

// compose connect string from raw config
func connectString(c config.MySQLConfig) string {

	return c.Username + ":" + c.Password + "@" + c.Host + "/" + c.Database

}
