package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/geowa4/base-go/components/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func configureGlobalLogging(level string) {
	zerolog.TimeFieldFormat = ""
	if logLevel, err := zerolog.ParseLevel(level); err != nil {
		zerolog.SetGlobalLevel(logLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func main() {
	configureGlobalLogging(os.Getenv("GOLOGLEVEL"))
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	cancelWebContext := server.NewWebServers(log.Logger)
	log.Info().Msg(fmt.Sprintf("Received signal %s; shutting down.", <-sigint))
	cancelWebContext()
}
