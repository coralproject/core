/*
Package config handles the loading and distribution of configuration

At Present:
- file: config.json
- schema:

{

	"Name": "Identity",

	"Database": {
		"username":  "",
		"password":  "",
		"host":      "",
		"port":     (int),
		"database": "",
		"adapter":  ""
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
const configFile = "config/config.json"

// Database config definition
type DatabaseConfig struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"pasword,string,omitempty"`
	Host     string `json:"host"`
	Port     int    `json:"port,string,omitempty"`
	Database string `json:"database"`
	Adapter  string `json:"adapter"`
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
func init() {
	config = readLocalFile(configFile)
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

// Get exposes the Databae config to others
func GetDatabase() DatabaseConfig {
	return config.Database
}
