/*
Package config handles the loading and distribution of configuration

At Present:
- file: config.json
- schema:

{

	"Name": "Identity",

	"MySQL": {
		"username": "",
		"password": "",
		"host": "",
		"port": (int),
		"database": ""
	},
	"Postgresql": {
		"username": "",
		"password": "",
		"host": "",
		"port": (int),
		"database": ""
	}

}

Future:
- load from s3
- load from db
- deeper schema for more configuratin goodness

*/

package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/coralproject/core/log"
)

// localFile is the localFile
const localFile = "./config.json"

// MySQL config definition
type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Adapter  string
}

// Top Level Config definition
type Config struct {
	Name     string
	Database DatabaseConfig
}

// Pointer to the master config record
var config *Config

// Load kicks off configuration loading
//  TODO: support multiple config sources
func Load() {
	config = readLocalFile(localFile)
}

func unmarshalConfig(content []byte) (*Config, error) {

	c := new(Config)

	if err := json.Unmarshal(content, &c); err != nil {
		return nil, err
	}

	return c, nil
}

// readLocalFile takes a filename (f),
// * reads the file (Fatal on failure)
// * unmarshalls (Fatal on failure)
// * returns the *Config
func readLocalFile(f string) *Config {

	content, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal("Unable to read config file '%s': %s", f, err)
	}

	c, err := unmarshalConfig(content)
	if err != nil {
		log.Fatal("Unable to parse JSON in config file '%s': %s", f, err)
	}

	return c

}

// Get exposes the top level config to others
func Get() *Config {
	return config
}

// Get exposes the MySQL config to others
func GetDatabase() DatabaseConfig {
	return config.Database
}
