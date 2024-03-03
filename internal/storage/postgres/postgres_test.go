package postgres_test

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/integtest"
	"github.com/stretchr/testify/require"
)

func setupTestPostgres(t *testing.T) (*pgxpool.Pool, func()) {
	t.Helper()
	if testing.Short() {
		t.Skip("skip long-running test in short mode")
	}

	pool, err := dockertest.NewPool("")
	require.NoError(t, err)
	require.NoError(t, pool.Client.Ping())

	res, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15.2-alpine",
		Env: []string{
			"POSTGRES_DB=integtest",
			"POSTGRES_USER=integtest",
			"POSTGRES_PASSWORD=integtest",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
		config.Tmpfs = map[string]string{"/var/lib/postgresql/data": "rw"}
	})
	require.NoError(t, err)
	require.NoError(t, res.Expire(uint((15 * time.Minute).Seconds())))

	host, port := integtest.GetHostPort(res, "5432/tcp")

	connURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("integtest", "integtest"),
		Host:   net.JoinHostPort(host, port),
		Path:   "integtest",
		RawQuery: url.Values{
			"sslmode":         []string{"disable"},
			"connect_timeout": []string{"5"},
		}.Encode(),
	}

	var db *pgxpool.Pool

	err = pool.Retry(func() (err error) {
		db, err = pgxpool.New(context.Background(), connURL.String())
		if err != nil {
			return err
		}
		return db.Ping(context.Background())
	})
	require.NoError(t, err)

	return db, func() {
		if err := pool.Purge(res); err != nil {
			t.Error(fmt.Errorf("purge resource: %w", err))
		}
	}
}
