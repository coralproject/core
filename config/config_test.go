package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var validMySQLConfig = `{

	"Name": "testName",

	"Database": {
		"username": "root",
		"password": "root",
		"host": "",
		"port": 3306,
		"database": "coral",
		"adapter": "mysql"
	}

}`

var validPostgreSQLConfig = `{

	"Name": "testName",

	"Database": {
		"username": "root",
		"password": "root",
		"host": "",
		"port": 3306,
		"database": "coral",
		"adapter": "postgresql"
	}

}`

var invalidConfig = `{

	"xName": "testName",

	"xMySQL": {
		"username": "root",
		"password": "root",
		"host": "",
		"port": 3306,
		"database": "coral",
		"adapter": "anything"
	}

}`

//func readLocalFile(f string) *Config {
func TestReadLocalFile(t *testing.T) {
	// create temporary local file
	var configFile string
	var c *Config
	c = readLocalFile(configFile)

	assert.Equal(t, c, validMySQLConfig, "they should be equal")

}

func TestUnmarshalConfig(t *testing.T) {

	c, _ := unmarshalConfig([]byte(validMySQLConfig))

	if c.Name != "testName" {
		t.Errorf("Name attribute did not properly unmarshall")
	}

}

// Test that the right database configuration is being returned
func TestGetDatabase(t *testing.T) {

	//assert.Equal(t, config.GetDatabase(), validDB, "they should be equal")

}

// Test response when configuration file is empty
func TestEmptyFile(t *testing.T) {

}

func TestMissingFile(t *testing.T) {
}

func TestConfigJSON(t *testing.T) {
}
