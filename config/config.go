package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/coralproject/core/log"
)

const (
	file = "./config.json"
)

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

type Config struct {
	Name string

	MySQL MySQLConfig
}

var (
	config *Config
)

func Load() {
	config = readLocalFile(file)
}

func readLocalFile(f string) *Config {

	c := new(Config)

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Unable to read config file '%s': %s", f, err)
	}

	if err = json.Unmarshal(content, &c); err != nil {
		log.Fatal("Unable to parse JSON in config file '%s': %s", f, err)
	}

	return c

}

func Get() *Config {
	return config
}
