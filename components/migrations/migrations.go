package migrations

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:generate go-bindata -pkg $GOPACKAGE -o embeds.go sql/...

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

//MigrateDatabase migrates the database to the latest version
//with the embedded scripts.
func MigrateDatabase(log zerolog.Logger, db *sql.DB) {
	createSchemaMigrationsTable(db)
	row := db.QueryRow("SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1")
	var version int
	err := row.Scan(&version)
	if err == sql.ErrNoRows {
		version = 0
	} else if err != nil {
		log.Fatal().Err(err).Msg("Error retrieving latest database version.")
	}
	log.Info().Msg(fmt.Sprintf("Current database version: %d", version))
	for nextVersion := version + 1; true; nextVersion++ {
		log.Info().Msg(fmt.Sprintf("Attempting to migrate database to version %d using sql/%04d.up.sql", nextVersion, nextVersion))
		migrationScript, err := Asset(fmt.Sprintf("sql/%04d.up.sql", nextVersion))
		if err != nil {
			log.Info().Msg(fmt.Sprintf("No database migration found for version %d.", nextVersion))
			break
		}
		_, err = db.Exec(string(migrationScript))
		if err != nil {
			log.Fatal().Err(err).Msg(fmt.Sprintf("Error migrating database to version %d.", nextVersion))
		}
		_, err = db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", nextVersion)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to insert into version table.")
		}
		log.Info().Msg(fmt.Sprintf("Applied database migration for version %d.", nextVersion))
	}
}
