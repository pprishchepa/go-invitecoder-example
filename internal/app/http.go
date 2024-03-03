package app

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/pprishchepa/go-invitecoder-example/internal/config"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func newHTTPServer(lc fx.Lifecycle, conf config.Config, handler http.Handler, logger zerolog.Logger) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.HTTP.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			l, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			logger.Info().Msgf("starting HTTP server on %s", srv.Addr)
			go func() {
				if err := srv.Serve(l); err != nil && !errors.Is(err, http.ErrServerClosed) {
					logger.Error().Err(err).Msg("HTTP server failed")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
