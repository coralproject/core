package database

import (
	"database/sql"
	"testing"
)

func init() {
}

/// create fake json file and fake db

var db *sql.DB

func TestConnectDB(t *testing.T) {
	db = Open()
}

func TestCloseDB(t *testing.T) {
	db.Close()
}

func TestQueryDB(t *testing.T) {
	db.Ping()
}
