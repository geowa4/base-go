package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"

	"github.com/geowa4/base-go/components/migrations"
	"github.com/geowa4/base-go/components/server"
	_ "github.com/lib/pq"
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

func createSchemaMigrationsTable(db *sql.DB) {
	const createSchemaMigrationsStatement = `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version smallint,
		created_at timestamp without time zone default (now() at time zone 'utc'),

		PRIMARY KEY (version)
	)
	`
	_, err := db.Exec(createSchemaMigrationsStatement)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating schema migrations table.")
	}
}

func main() {
	configureGlobalLogging(os.Getenv("GOLOGLEVEL"))
	log.Info().Msg("Initializing application.")
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database.")
	}
	defer db.Close()
	log.Info().Msg("Connected to the database.")
	createSchemaMigrationsTable(db)
	migrations.MigrateDatabase(db)
	cancelWebContext := server.NewWebServers(log.Logger)
	defer cancelWebContext()
	log.Info().Msg("Application ready.")
	log.Info().Msg(fmt.Sprintf("Received signal %s; shutting down.", <-sigint))
}
