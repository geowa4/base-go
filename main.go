package main

import (
	"database/sql"
	"os"
	"os/signal"

	"github.com/jmoiron/sqlx"

	"github.com/geowa4/base-go/components/config"
	"github.com/geowa4/base-go/components/migrations"
	"github.com/geowa4/base-go/components/server"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

//serviceName is used as a prefix to environment variables and put in logs for easy filtering.
const serviceName = "base_go"

func configureGlobalLogging() zerolog.Logger {
	level := viper.GetString("log_level")
	zerolog.TimeFieldFormat = ""
	if logLevel, err := zerolog.ParseLevel(level); err != nil {
		zerolog.SetGlobalLevel(logLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	return log.With().Str("service", serviceName).Logger()
}

func main() {
	config.ReadConfig(serviceName)
	logger := configureGlobalLogging()
	logger.Info().Msg("Initializing application.")
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	db, err := sql.Open("postgres", config.DatabaseConnectionString())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to the database.")
	}
	defer db.Close()
	logger.Info().Msg("Configured connection to the database.")
	migrations.MigrateDatabase(logger, db)
	cancelWebContext := server.NewWebServers(logger, sqlx.NewDb(db, "postgres"))
	defer cancelWebContext()
	logger.Info().Msg("Application ready.")
	logger.Info().Msgf("Received signal %s; shutting down.", <-sigint)
}
