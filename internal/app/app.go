package app

import (
	"net/http"

	"github.com/pprishchepa/go-invitecoder-example/internal/config"
	httpctrl "github.com/pprishchepa/go-invitecoder-example/internal/controller/http"
	httpv1 "github.com/pprishchepa/go-invitecoder-example/internal/controller/http/v1"
	"github.com/pprishchepa/go-invitecoder-example/internal/pkg/fxlog"
	"github.com/pprishchepa/go-invitecoder-example/internal/storage/postgres"
	"github.com/pprishchepa/go-invitecoder-example/internal/usecase"
	"github.com/rs/zerolog"
	"go.uber.org/automaxprocs/maxprocs"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			newLogger,
		),
		fx.Provide(
			newDBStatsClient,
			newDBUserClient,
			newStatsStorage,
			newUserStorage,
			usecase.NewInviteService,
			httpv1.NewInvitesRoutes,
			httpctrl.NewRouter,
			newHTTPServer,
			func(v *usecase.InviteService) httpv1.InviteService { return v },
			func(v *postgres.StatsStorage) usecase.StatsStorage { return v },
			func(v *postgres.UserStorage) usecase.UserStorage { return v },
		),
		fx.WithLogger(newZerologAdapter),
		fx.Invoke(automaxprocs),
		fx.Invoke(migrate),
		fx.Invoke(func(*http.Server) {}),
	)
}

func automaxprocs(logger zerolog.Logger) error {
	_, err := maxprocs.Set(maxprocs.Logger(func(s string, i ...interface{}) {
		logger.Debug().Msgf(s, i...)
	}))
	return err
}

func newZerologAdapter(logger zerolog.Logger) fxevent.Logger {
	return fxlog.NewZerologAdapter(logger)
}
