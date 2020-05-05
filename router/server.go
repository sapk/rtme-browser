package router

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
)

const contextTimeout = 15*time.Second

//StartServer start the server
func StartServer(u url.URL, allowCORS bool) {
	r := generateRouter(allowCORS)
	// Listen and Server (endlessly) on http://0.0.0.0:3000
	srv := &http.Server{
		Addr:    u.Host,
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Fail to listen on %s: %s", u.String(), err)
		}
	}()

	log.Info().Msgf("Listening on: %s", u.String())

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 15 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("Server Shutdown: %v", err)
	}
}
