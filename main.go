package main

//go:generate swagger generate spec -o ./assets/swagger/swagger.v1.json
//go:generate esc -pkg=swagger -prefix=assets/swagger -ignore='.*.map' -o public/swagger/swagger.go assets/swagger/
//go:generate esc -pkg=ui -prefix=assets/ui/dist -ignore='.*.map' -o public/ui/ui.go assets/ui/dist/

import (
	"flag"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/pkg/browser"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/sapk/go-genesys/db"

	"github.com/sapk/rtme-browser/router"
)

func main() {
	//TODO use cobra for validation and use env and file config
	debug := flag.Bool("v", false, "sets log level to debug")
	listenAddr := flag.String("addr", ":3000", "listening address")
	dbType := flag.String("db-type", "mssql", "database address")
	dbRTMEURL := flag.String("db-rtme", "server=localhost;user id=sa;password=g3n3sys;database=RTME", "RTME database url definition (xorm definition)")
	dbCFGURL := flag.String("db-cfg", "server=localhost;user id=sa;password=g3n3sys;database=CFG", "CFG database url definition (xorm definition)")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: colorable.NewColorableStdout(),
	})

	u := parseAddr(*listenAddr)
	quit := startServer(u, *dbType, *dbRTMEURL, *dbCFGURL)
	go startBrowser(u)
	<-quit //Wait for quit
}

func startServer(u url.URL, dbType, dbRTMEURL, dbCFGURL string) chan bool {
	quit := make(chan bool)
	//DB connexion
	rtmeDB, err := db.NewDBFromURL(dbType, dbRTMEURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to RTME database")
		quit <- true
		return quit
	}
	cfgDB, err := db.NewDBFromURL(dbType, dbCFGURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to CFG database")
		quit <- true
		return quit
	}

	//Setup server
	gin.SetMode(gin.ReleaseMode)
	go func() {
		router.StartServer(u, rtmeDB, cfgDB)
		quit <- true
	}()
	return quit
}

func startBrowser(u url.URL) {
	browser.OpenURL(u.String())
}

func parseAddr(addr string) url.URL {
	u := url.URL{
		Scheme: "http",
		Host:   addr,
	}
	if strings.HasPrefix(u.Host, ":") {
		u.Host = "0.0.0.0" + u.Host
	}
	return u
}
