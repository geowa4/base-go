package migrations

import (
	"database/sql"

	"github.com/geowa4/base-go/components/migrations/internal/assets"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/rs/zerolog"
)

//go:generate go-bindata -pkg assets -o internal/assets/embeds.go -prefix sql/ sql/...

func newSourceDriver() (source.Driver, error) {
	s := bindata.Resource(
		assets.AssetNames(),
		func(name string) ([]byte, error) {
			return assets.Asset(name)
		},
	)

	return bindata.WithInstance(s)
}

//MigrateDatabase migrates the database to the latest version
//with the embedded scripts.
func MigrateDatabase(log zerolog.Logger, db *sql.DB) {
	var (
		err error
		m   *migrate.Migrate
	)

	sourceDriver, err := newSourceDriver()
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating bindata source driver.")
	}

	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating database driver.")
	}
	m, err = migrate.NewWithInstance("go-bindata", sourceDriver, "postgres", dbDriver)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating migration from driver.")
	}
	runMigration(log, m)
}

func runMigration(log zerolog.Logger, m *migrate.Migrate) {
	err := m.Up()
	if err == migrate.ErrNoChange {
		log.Info().Msg("No database migrations to apply.")
	} else if err != nil {
		log.Fatal().Err(err).Msg("Error running migration.")
	} else {
		log.Info().Msg("Schema migrations applied successfully.")
	}
}
