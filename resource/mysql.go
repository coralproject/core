package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/coralproject/core/config"
	"github.com/coralproject/core/log"
)

var Db *sql.DB

func Connect() *sql.DB {

	// get the config
	c := config.GetMySQL()

	// build the connect string for MySQL
	s := connectString(c)

	// connect!
	Db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal("Could not connect to MySQL with", s, err)
	}

	return Db

}

func connectString(c config.MySQLConfig) string {

	return c.Username + ":" + c.Password + "@" + c.Host + "/" + c.Database

}
