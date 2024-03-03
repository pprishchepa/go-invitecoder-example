package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
)

type InviteStatsStorage struct {
	db *pgxpool.Pool
}

func (s InviteStatsStorage) GetValues(ctx context.Context) (map[string]int, error) {
	rows, err := s.db.Query(ctx, `SELECT code, accepted FROM invite_stats`)
	if err != nil {
		return nil, fmt.Errorf("exec sql: %w", err)
	}
	defer rows.Close()

	values := make(map[string]int, 16)
	for rows.Next() {
		var code string
		var accepted int
		if err := rows.Scan(&code, &accepted); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		values[code] = accepted
	}

	return values, nil
}

func (s InviteStatsStorage) IncByCode(ctx context.Context, code string, maxVal int) error {
	sql := `
		INSERT INTO invite_stats (code, accepted)
		VALUES ($1, 1)
		ON CONFLICT (code) DO UPDATE SET accepted = invite_stats.accepted + 1
		WHERE invite_stats.accepted < $2
	`

	cmd, err := s.db.Exec(ctx, sql, code, maxVal)
	if err != nil {
		return fmt.Errorf("exec sql: %w", err)
	}

	if cmd.RowsAffected() == 0 {
		return entity.ErrNotAvailable
	}

	return nil
}

func NewInviteStatsStorage(db *pgxpool.Pool) *InviteStatsStorage {
	return &InviteStatsStorage{db: db}
}
