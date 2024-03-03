package app

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pprishchepa/go-invitecoder-example/internal/config"
	"github.com/pprishchepa/go-invitecoder-example/migrations/stats"
	"golang.org/x/sync/errgroup"
)

func performMigrations(conf config.Config) error {
	var eg errgroup.Group

	eg.Go(func() error {
		if err := migratePgCounters(conf); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate counters: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		if err := migratePgEmails01(conf); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate emails01: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		if err := migratePgEmails02(conf); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate emails02: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		if err := migratePgEmails03(conf); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate emails03: %w", err)
		}
		return nil
	})

	return eg.Wait()
}

func migratePgCounters(conf config.Config) error {
	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.PgCounters.User, conf.PgCounters.Password),
		Host:   net.JoinHostPort(conf.PgCounters.Host, strconv.Itoa(conf.PgCounters.Port)),
		Path:   conf.PgCounters.Database,
		RawQuery: url.Values{
			"sslmode": []string{conf.PgCounters.SSLMode},
		}.Encode(),
	}

	return stats.Migrate(connURL.String())
}

func migratePgEmails01(conf config.Config) error {
	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.PgEmails01.User, conf.PgEmails01.Password),
		Host:   net.JoinHostPort(conf.PgEmails01.Host, strconv.Itoa(conf.PgEmails01.Port)),
		Path:   conf.PgEmails01.Database,
		RawQuery: url.Values{
			"sslmode": []string{conf.PgEmails01.SSLMode},
		}.Encode(),
	}

	return invitee.Migrate(connURL.String())
}

func migratePgEmails02(conf config.Config) error {
	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.PgEmails02.User, conf.PgEmails02.Password),
		Host:   net.JoinHostPort(conf.PgEmails02.Host, strconv.Itoa(conf.PgEmails02.Port)),
		Path:   conf.PgEmails02.Database,
		RawQuery: url.Values{
			"sslmode": []string{conf.PgEmails02.SSLMode},
		}.Encode(),
	}

	return invitee.Migrate(connURL.String())
}

func migratePgEmails03(conf config.Config) error {
	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(conf.PgEmails03.User, conf.PgEmails03.Password),
		Host:   net.JoinHostPort(conf.PgEmails03.Host, strconv.Itoa(conf.PgEmails03.Port)),
		Path:   conf.PgEmails03.Database,
		RawQuery: url.Values{
			"sslmode": []string{conf.PgEmails03.SSLMode},
		}.Encode(),
	}

	return invitee.Migrate(connURL.String())
}
