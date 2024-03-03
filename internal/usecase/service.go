package usecase

import (
	"context"
	"fmt"

	"github.com/pprishchepa/go-invitecoder-example/internal/config"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
)

//go:generate go run go.uber.org/mock/mockgen -source=service.go -destination=service_mock_test.go -package=usecase_test

type (
	StatsStorage interface {
		IncByCode(ctx context.Context, code string, maxVal int) error
	}
	UserStorage interface {
		SaveUser(ctx context.Context, user entity.InvitedUser) error
	}
)

type InviteService struct {
	conf  config.Config
	stats StatsStorage
	users UserStorage
}

func (s *InviteService) AcceptInvite(ctx context.Context, user entity.InvitedUser) error {
	if err := s.stats.IncByCode(ctx, user.InvitedVia, s.conf.MaxUsersPerCode); err != nil {
		return fmt.Errorf("inc by code: %w", err)
	}

	if err := s.users.SaveUser(ctx, user); err != nil {
		return fmt.Errorf("save user: %w", err)
	}

	return nil
}

func NewInviteService(conf config.Config, stats StatsStorage, users UserStorage) *InviteService {
	return &InviteService{conf: conf, stats: stats, users: users}
}
