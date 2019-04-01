package config

import (
	"fmt"

	"github.com/sapk/go-genesys/db"
)

//Config hold the config and db access
type Config struct {
	RTME *db.DB
	CFG  *db.DB
}

var (
	config *Config
)

//TODO manage lock

//IsNotInit state if the config is not init
func IsNotInit() bool {
	return config == nil
}

//Setup start the db connexion
func Setup(dbType, dbRTMEURL, dbCFGURL string) error {
	rtmeDB, err := db.NewDBFromURL(dbType, dbRTMEURL)
	if err != nil {
		return err
	}
	cfgDB, err := db.NewDBFromURL(dbType, dbCFGURL)
	if err != nil {
		return err
	}
	config = &Config{
		RTME: rtmeDB,
		CFG:  cfgDB,
	}
	return nil
}

//Get give access to the config object
func Get() (*Config, error) {
	if IsNotInit() {
		return nil, fmt.Errorf("Config is not set")
	}
	return config, nil
}
