package usecase_test

import (
	"context"
	"testing"

	"github.com/pprishchepa/go-invitecoder-example/internal/config"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/pprishchepa/go-invitecoder-example/internal/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestInviteService(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	conf := config.Config{
		MaxUsersPerCode: 1000,
	}

	stats := NewMockStatsStorage(mockCtrl)
	users := NewMockUserStorage(mockCtrl)

	svc := usecase.NewInviteService(conf, stats, users)

	stats.EXPECT().IncByCode(ctx, "twitter", 1000).Return(nil)
	users.EXPECT().SaveUser(ctx, entity.InvitedUser{Email: "foo@example.com", InvitedVia: "twitter"}).Return(nil)
	require.NoError(t, svc.AcceptInvite(ctx, entity.InvitedUser{Email: "foo@example.com", InvitedVia: "twitter"}))
}
