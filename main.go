package main

//go:generate go run -mod=vendor github.com/go-swagger/go-swagger/cmd/swagger generate spec -o ./assets/swagger/swagger.v1.json

import (
	"flag"
	"net"
	"net/url"
	"os"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/pkg/browser"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zserge/lorca"

	"github.com/sapk/rtme-browser/router"
)

func main() {
	//TODO use cobra for validation and use env and file config
	debug := flag.Bool("v", false, "sets log level to debug")
	noWebview := flag.Bool("no-webview", false, "de-activate the start of the webview")
	browser := flag.Bool("browser", false, "de-activate the start of the browser")
	listenAddr := flag.String("addr", "localhost:0", "listening address")
	allowCORS := flag.Bool("cors", false, "allow cors")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: colorable.NewColorableStdout(),
	})

	u := parseAddr(*listenAddr)
	quit := startServer(u, *allowCORS)
	if !*noWebview {
		go startWebview(u)
	}
	if *browser {
		go startBrowser(u)
	}
	<-quit //Wait for quit
}

func startServer(u url.URL, allowCORS bool) chan bool {
	quit := make(chan bool)

	//Setup server
	gin.SetMode(gin.ReleaseMode)
	go func() {
		router.StartServer(u, allowCORS)
		quit <- true
	}()
	return quit
}

func startBrowser(u url.URL) {
	err := browser.OpenURL(u.String())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the browser")
	}
}

func startWebview(u url.URL) {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 800, 600, args...)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the chrome webview")
	}
	err = ui.Load(u.String())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to setup the chrome webview")
	}
	// Wait for the browser window to be closed
	<-ui.Done()
	os.Exit(0) //Stop all
}

func parseAddr(addr string) url.URL {
	u := url.URL{
		Scheme: "http",
		Host:   addr,
	}
	if strings.HasPrefix(u.Host, ":") {
		u.Host = "0.0.0.0" + u.Host
	}
	if u.Port() == "0" {
		u.Host = getHostWithFreePort(u.Hostname())
	}
	return u
}

func getHostWithFreePort(host string) string {
	ln, err := net.Listen("tcp", host+":0")

	if err != nil {
		log.Fatal().Err(err).Msg("Failed open a tcp listening port")
	}

	err = ln.Close()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to free a tcp listening port")
	}

	return ln.Addr().String()
}
