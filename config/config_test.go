package config

import (
	"testing"
)

var validConfig = `{

	"Name": "testName",

	"MySQL": {
		"username": "root",
		"password": "root",
		"host": "",
		"port": 3306,
		"database": "coral"
	}

}`

var invalidConfig = `{

	"xName": "testName",

	"xMySQL": {
		"username": "root",
		"password": "root",
		"host": "",
		"port": 3306,
		"database": "coral"
	}

}`

/*
var mysqlConfig = MySQLConfig{
	Username: "testName",
	Password: "testPass",
	Host:     "testHost",
	Port:     42,
	Database: "testBase",
}

var config = Config{
	Name:  "TestApp",
	MySQL: mysqlConfig,
}
*/

func TestUnmarshalConfig(t *testing.T) {

	c, _ := unmarshalConfig([]byte(validConfig))

	if c.Name != "testName" {
		t.Errorf("Name attribute did not properly unmarshall")
	}

}
