package postgres_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/pgxsharded"
	"github.com/pprishchepa/go-invitecoder-example/internal/storage/postgres"
	"github.com/pprishchepa/go-invitecoder-example/migrations/user"
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

	require.NoError(t, user.Migrate(db1.Config().ConnString()))
	require.NoError(t, user.Migrate(db2.Config().ConnString()))
	require.NoError(t, user.Migrate(db3.Config().ConnString()))

	ctx := context.Background()
	cluster := pgxsharded.NewCluster([]*pgxpool.Pool{db1, db2, db3})
	storage := postgres.NewInvitedUserStorage(cluster)

	err := storage.SaveUser(ctx, entity.InvitedUser{Email: "foo@example.com", InvitedVia: "twitter"})
	require.NoError(t, err)

	err = storage.SaveUser(ctx, entity.InvitedUser{Email: "bar@example.com", InvitedVia: "twitter"})
	require.NoError(t, err)

	err = storage.SaveUser(ctx, entity.InvitedUser{Email: "baz@example.com", InvitedVia: "twitter"})
	require.NoError(t, err)

	err = storage.SaveUser(ctx, entity.InvitedUser{Email: "foo@example.com", InvitedVia: "fb"})
	require.ErrorIs(t, entity.ErrAlreadyExists, err)
}
