package pgxmigrator

import (
	"embed"
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog"
)

type Migrator struct {
	logger zerolog.Logger
}

func NewMigrator(logger zerolog.Logger) *Migrator {
	return &Migrator{
		logger: logger.With().Str("logger", "Migrator").Logger(),
	}
}

func (m Migrator) Up(db *pgxpool.Pool, fs embed.FS) error {
	hostname := net.JoinHostPort(db.Config().ConnConfig.Host, strconv.Itoa(int(db.Config().ConnConfig.Port)))

	mg, err := m.newInstance(db, fs)
	if err != nil {
		return fmt.Errorf("init migrator: %s: %w", hostname, err)
	}

	m.logger.Info().Str("dbhost", hostname).Msg("migrating...")

	if err := mg.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) || errors.Is(err, migrate.ErrNilVersion) {
			m.logger.Info().Str("dbhost", hostname).Msg(err.Error())
			return nil
		}
		m.logger.Err(err).Str("dbhost", hostname).Msg("could not migrate")
		return fmt.Errorf("migrate: %s: %w", hostname, err)
	}

	m.logger.Info().Str("dbhost", hostname).Msg("migration succeeded")
	return nil
}

func (m Migrator) newInstance(db *pgxpool.Pool, fs embed.FS) (*migrate.Migrate, error) {
	source, err := iofs.New(fs, ".")
	if err != nil {
		return nil, fmt.Errorf("create source: %w", err)
	}

	stdDB := stdlib.OpenDBFromPool(db)
	defer func() { _ = stdDB.Close() }()

	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.InitialInterval = 2 * time.Second
	expBackoff.MaxInterval = 10 * time.Second
	expBackoff.MaxElapsedTime = 5 * time.Minute

	var driver database.Driver
	err = backoff.Retry(func() error {
		var err error
		driver, err = postgres.WithInstance(stdDB, &postgres.Config{})
		if err != nil {
			m.logger.Warn().Err(err).Msg("database connection issue, retrying...")
			return err
		}
		return nil
	}, expBackoff)
	if err != nil {
		return nil, fmt.Errorf("create driver: %w", err)
	}

	return migrate.NewWithInstance("iofs", source, db.Config().ConnConfig.Database, driver)
}
