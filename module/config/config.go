package config

import (
	"fmt"

	"github.com/sapk/go-genesys/db"
)

//Config hold the config and db access.
type Config struct {
	DB *db.DB
}

var (
	config *Config
)

//TODO manage lock

//IsNotInit state if the config is not init.
func IsNotInit() bool {
	return config == nil
}

//Setup start the db connexion.
func Setup(dbType, dbRTMEURL, dbCFGURL string) error {
	d := db.NewDBFromURL(dbType, dbCFGURL, dbRTMEURL)
	e, err := d.CFG()
	if err != nil {
		return err
	}
	err = e.Ping()
	if err != nil {
		return err
	}
	e, err = d.RTME()
	if err != nil {
		return err
	}
	err = e.Ping()
	if err != nil {
		return err
	}
	config = &Config{
		DB: d,
	}
	return nil
}

//Get give access to the config object.
func Get() (*Config, error) {
	if IsNotInit() {
		return nil, fmt.Errorf("config is not set")
	}
	return config, nil
}
