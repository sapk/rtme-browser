package config

import "github.com/sapk/go-genesys/db"

var (
	dBRTME *db.DB
	dBCFG *db.DB
)

//TODO manage lock

//IsNotInit state if the config is not init
func IsNotInit() bool  {
	return dBRTME == nil || dBCFG == nil
}

//Setup start the db connexion
func Setup(dbType, dbRTMEURL, dbCFGURL string) error  {
	rtmeDB, err := db.NewDBFromURL(dbType, dbRTMEURL)
	if err != nil {
		return err
	}
	cfgDB, err := db.NewDBFromURL(dbType, dbCFGURL)
	if err != nil {
		return err
	}
	dBCFG = cfgDB
	dBRTME = rtmeDB
	return nil
}