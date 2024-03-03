package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/pprishchepa/go-invitecoder-example/internal/config"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/rs/zerolog"
)

//go:generate go run go.uber.org/mock/mockgen -source=service.go -destination=service_mock_test.go -package=usecase_test

type (
	StatsStorage interface {
		IncByCode(ctx context.Context, code string, maxVal int) error
		DecByCode(ctx context.Context, code string) error
	}
	UserStorage interface {
		SaveUser(ctx context.Context, user entity.InvitedUser) error
	}
)

type InviteService struct {
	conf   config.Config
	stats  StatsStorage
	users  UserStorage
	logger zerolog.Logger
}

func NewInviteService(conf config.Config, stats StatsStorage, users UserStorage, logger zerolog.Logger) *InviteService {
	return &InviteService{
		conf:   conf,
		stats:  stats,
		users:  users,
		logger: logger,
	}
}

func (s *InviteService) AcceptInvite(ctx context.Context, user entity.InvitedUser) error {
	if err := s.stats.IncByCode(ctx, user.InvitedVia, s.conf.MaxUsersPerCode); err != nil {
		return fmt.Errorf("inc by code: %w", err)
	}

	if err := s.users.SaveUser(ctx, user); err != nil {
		if errors.Is(err, entity.ErrAlreadyExists) {
			if err := s.stats.DecByCode(ctx, user.InvitedVia); err != nil {
				s.logger.Err(err).Msg("could not rollback stats increment")
			}
		}
		return fmt.Errorf("save user: %w", err)
	}

	return nil
}
