/*

Package database

Get configuration for the database and makes the connection.

Status: Stub

TO DO

* Ping to check connection is not convincing
* Select an ORM (maybe gorp)

*/

package database

import (
	_ "github.com/go-sql-driver/mysql" // Blank import because we are using psql or mysql depending on configuration
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Blank import because we are using psql or mysql depending on configuration
)

type Driver interface {
	Open() (*gorm.DB, error)
	Close(db *gorm.DB) error
}
