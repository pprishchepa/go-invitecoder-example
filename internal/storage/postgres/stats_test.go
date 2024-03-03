package postgres_test

import (
	"context"
	"testing"

	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/pprishchepa/go-invitecoder-example/internal/storage/postgres"
	"github.com/pprishchepa/go-invitecoder-example/migrations/stats"
	"github.com/stretchr/testify/require"
)

func TestInviteStatsStorage(t *testing.T) {

	t.Parallel()

	db, teardownDB := setupTestPostgres(t)
	defer func() { teardownDB() }()

	require.NoError(t, stats.Migrate(db.Config().ConnString()))

	const maxVal = 2
	ctx := context.Background()
	storage := postgres.NewInviteStatsStorage(db)

	require.NoError(t, storage.IncByCode(ctx, "foo", maxVal))
	require.NoError(t, storage.IncByCode(ctx, "bar", maxVal))
	require.NoError(t, storage.IncByCode(ctx, "foo", maxVal))
	require.NoError(t, storage.IncByCode(ctx, "baz", maxVal))
	require.ErrorIs(t, storage.IncByCode(ctx, "foo", maxVal), entity.ErrNotAvailable)

	values, err := storage.GetValues(ctx)
	require.NoError(t, err)
	require.Equal(t, map[string]int{
		"foo": 2,
		"bar": 1,
		"baz": 1,
	}, values)
}
