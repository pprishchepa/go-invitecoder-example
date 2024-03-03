package postgres_test

import (
	"context"
	"os"
	"testing"

	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/pgxmigrator"
	"github.com/pprishchepa/go-invitecoder-example/internal/storage/postgres"
	"github.com/pprishchepa/go-invitecoder-example/migrations/dbstats"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestStatsStorage(t *testing.T) {
	t.Parallel()

	db, teardownDB := setupTestPostgres(t)
	defer func() { teardownDB() }()

	migrator := pgxmigrator.NewMigrator(zerolog.New(os.Stdout))
	require.NoError(t, migrator.Up(db, dbstats.FS))

	const maxVal = 2
	ctx := context.Background()
	storage := postgres.NewStatsStorage(db)

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
