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

	"github.com/sapk-fork/webview"

	"github.com/sapk/rtme-browser/router"
)

func main() {
	//TODO use cobra for validation and use env and file config
	debug := flag.Bool("v", false, "sets log level to debug")
	noWebview := flag.Bool("no-webview", false, "de-activate the start of the webview")
	browser := flag.Bool("browser", false, "de-activate the start of the browser")
	listenAddr := flag.String("addr", "localhost:3000", "listening address")
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
	browser.OpenURL(u.String())
}

func startWebview(u url.URL) {
	//webview.Open("rtme-browser", u.String(), 800, 600, true)

	w := webview.New(true)
	w.Navigate(u.String())
	w.SetTitle("rtme-browser")
	//TODO set_size
	/*
		w.Dispatch(func() {
			println("Hello dispatch")
		})
	*/
	w.Run()
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
