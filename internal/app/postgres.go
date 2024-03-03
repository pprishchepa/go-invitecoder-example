package app

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pprishchepa/go-invitecoder-example/internal/config"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/pgxcluster"
	"go.uber.org/fx"
)

type (
	dbstatsPool struct {
		*pgxpool.Pool
	}
	dbusersCluster struct {
		*pgxcluster.Cluster
	}
)

func newDBStatsClient(lc fx.Lifecycle, conf config.Config) (*dbstatsPool, error) {
	db, err := newPostgresClient(conf.DBStats)
	if err != nil {
		hostname := net.JoinHostPort(conf.DBStats.Host, strconv.Itoa(conf.DBStats.Port))
		return nil, fmt.Errorf("new dbstats client (%s): %w", hostname, err)
	}

	lc.Append(fx.StopHook(func() {
		db.Close()
	}))

	return &dbstatsPool{db}, nil
}

func newDBUserClient(lc fx.Lifecycle, conf config.Config) (*dbusersCluster, error) {
	const dbusersShards = 3

	shards := make([]*pgxpool.Pool, 0, dbusersShards)

	for _, shardConf := range []config.Postgres{
		conf.DBUsers01,
		conf.DBUsers02,
		conf.DBUsers03,
	} {
		db, err := newPostgresClient(shardConf)
		if err != nil {
			hostname := net.JoinHostPort(shardConf.Host, strconv.Itoa(shardConf.Port))
			return nil, fmt.Errorf("new dbusers client (%s): %w", hostname, err)
		}
		lc.Append(fx.StopHook(func() {
			db.Close()
		}))
		shards = append(shards, db)
	}

	return &dbusersCluster{pgxcluster.NewCluster(shards)}, nil
}

func newPostgresClient(conf config.Postgres) (*pgxpool.Pool, error) {
	connString := strings.Join([]string{
		"user=" + conf.User,
		"password=" + conf.Password,
		"dbname=" + conf.Database,
		"host=" + conf.Host,
		"port=" + fmt.Sprintf("%d", conf.Port),
		"sslmode=" + conf.SSLMode,
		"connect_timeout=" + fmt.Sprintf("%d", conf.ConnTimeout),
		"pool_max_conns=" + fmt.Sprintf("%d", conf.MaxConn),
	}, " ")

	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("init pgxpool: %w", err)
	}

	return db, nil
}
