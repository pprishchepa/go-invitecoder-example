package app

import (
	"github.com/pprishchepa/go-invitecoder-example/internal/storage/postgres"
)

func newStatsStorage(db *dbstatsPool) (*postgres.StatsStorage, error) {
	return postgres.NewStatsStorage(db.Pool), nil
}

func newUserStorage(db *dbusersCluster) (*postgres.UserStorage, error) {
	return postgres.NewUserStorage(db.Cluster), nil
}
