package resource

import (
	"database/sql"
	"log"

	"github.com/coralproject/core/config"
)

// db points to the db
var Db *sql.DB

// Open the db connection using coralproject/core/config creds
func Open() *sql.DB {

	// get the config
	c := config.GetDatabase()

	// build the connect string for MySQL
	var s string
	if c.Adapter == "mysql" {
		s = getMySQLConnection(c)
	} else if c.Adapter == "postgresql" {
		s = getPostgresqlConnection(c)
	}

	// open, or so it appears, but not yet
	db, err := sql.Open(c.Adapter, s)
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

func getMySQLConnection(c config.DatabaseConfig) string {
	return c.Username + ":" + c.Password + "@" + c.Host + "/" + c.Database
}

func getPostgresqlConnection(c config.DatabaseConfig) string {
	return "user=" + c.Username + ":" + c.Password + " dbname=" + c.Database + "sslmode=disable"
}
