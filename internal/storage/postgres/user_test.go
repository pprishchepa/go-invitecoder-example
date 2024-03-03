package postgres_test

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/pgxcluster"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/pgxmigrator"
	"github.com/pprishchepa/go-invitecoder-example/internal/storage/postgres"
	"github.com/pprishchepa/go-invitecoder-example/migrations/dbusers"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestUserStorage(t *testing.T) {
	t.Parallel()

	db1, teardownDB1 := setupTestPostgres(t)
	db2, teardownDB2 := setupTestPostgres(t)
	db3, teardownDB3 := setupTestPostgres(t)
	defer func() {
		teardownDB1()
		teardownDB2()
		teardownDB3()
	}()

	migrator := pgxmigrator.NewMigrator(zerolog.New(os.Stdout))
	require.NoError(t, migrator.Up(db1, dbusers.FS))
	require.NoError(t, migrator.Up(db2, dbusers.FS))
	require.NoError(t, migrator.Up(db3, dbusers.FS))

	ctx := context.Background()
	cluster := pgxcluster.NewCluster([]*pgxpool.Pool{db1, db2, db3})
	storage := postgres.NewUserStorage(cluster)

	err := storage.SaveUser(ctx, entity.InvitedUser{Email: "foo@example.com", InvitedVia: "twitter"})
	require.NoError(t, err)

	err = storage.SaveUser(ctx, entity.InvitedUser{Email: "bar@example.com", InvitedVia: "twitter"})
	require.NoError(t, err)

	err = storage.SaveUser(ctx, entity.InvitedUser{Email: "baz@example.com", InvitedVia: "twitter"})
	require.NoError(t, err)

	err = storage.SaveUser(ctx, entity.InvitedUser{Email: "foo@example.com", InvitedVia: "fb"})
	require.ErrorIs(t, entity.ErrAlreadyExists, err)
}
