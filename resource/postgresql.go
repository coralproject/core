package resource

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/coralproject/core/config"
	"github.com/coralproject/core/log"
)

//type XXX interface {}

// db points to the postgresl db
var db *sql.DB

// Open the db connection using coralproject/core/config credentials
func Open() *sql.DB {

	configPostgresql := config.GetPostgresql()
	connectString := connectString(configPostgresql)

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal("Could not connect to Posgresql with", s, err)
	}

	return db
}

func Close() {
	db.Close()
}

// compose connect string from raw config
func connectString(c config.PostgresqlConfig) string {
	return "user=" + c.Username + ":" + c.Password + " dbname=" + c.Database + "sslmode=disable"
}
