package app

import (
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/pgxmigrator"
	"github.com/pprishchepa/go-invitecoder-example/migrations/dbstats"
	"github.com/pprishchepa/go-invitecoder-example/migrations/dbusers"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

func migrate(statsDB *dbstatsPool, usersDB *dbusersCluster, logger zerolog.Logger) error {
	mg := pgxmigrator.NewMigrator(logger)

	var eg errgroup.Group

	eg.Go(func() error {
		return mg.Up(statsDB.Pool, dbstats.FS)
	})

	for i := 0; i < usersDB.Size(); i++ {
		db, err := usersDB.GetShard(i)
		if err != nil {
			return err
		}
		eg.Go(func() error {
			return mg.Up(db, dbusers.FS)
		})
	}

	return eg.Wait()
}
