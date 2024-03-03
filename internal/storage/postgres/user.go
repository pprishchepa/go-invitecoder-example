package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/hashing"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/pgxsharded"
)

// see https://www.postgresql.org/docs/current/errcodes-appendix.html
const errorCodeUniqueViolation = "23505"

type InviteUserStorage struct {
	cluster *pgxsharded.Cluster
}

func (s InviteUserStorage) SaveUser(ctx context.Context, user entity.InvitedUser) error {
	shardID := hashing.HashStringKey(user.Email, s.cluster.Size())

	db, err := s.cluster.GetShard(int(shardID))
	if err != nil {
		return fmt.Errorf("get shard: %w", err)
	}

	_, err = db.Exec(ctx, "INSERT INTO invite_user (email, invited_via) VALUES ($1, $2)", user.Email, user.InvitedVia)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == errorCodeUniqueViolation {
			return entity.ErrAlreadyExists
		}
		return fmt.Errorf("exec sql: %w", err)
	}

	return nil
}

func NewInvitedUserStorage(cluster *pgxsharded.Cluster) *InviteUserStorage {
	return &InviteUserStorage{cluster: cluster}
}
