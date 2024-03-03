package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pprishchepa/go-invitecoder-example/internal/controller/http/v1/model"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/rs/zerolog"
)

//go:generate go run go.uber.org/mock/mockgen -source=invites.go -destination=invites_mock_test.go -package=v1_test

type InviteService interface {
	AcceptInvite(ctx context.Context, user entity.InvitedUser) error
}

type InvitesRoutes struct {
	service InviteService
	logger  zerolog.Logger
}

func NewInvitesRoutes(service InviteService, logger zerolog.Logger) *InvitesRoutes {
	return &InvitesRoutes{
		service: service,
		logger:  logger.With().Str("logger", "InvitesRoutes").Logger(),
	}
}

func (r *InvitesRoutes) RegisterRoutes(e *gin.RouterGroup) {
	e.POST("/invitation", r.acceptInvite)
}

func (r *InvitesRoutes) acceptInvite(c *gin.Context) {
	var m model.AcceptInviteRequest
	if err := c.ShouldBindJSON(&m); err != nil {
		r.logger.Debug().Err(err).Msg("invalid request")
		c.Status(http.StatusBadRequest)
		return
	}

	err := r.service.AcceptInvite(c.Request.Context(), entity.InvitedUser{
		Email:      m.Email,
		InvitedVia: m.Code,
	})
	if err != nil {
		r.logger.Debug().Err(err).Str("email", m.Email).Str("code", m.Code).
			Msg("failed to accept invite")
		switch {
		case errors.Is(err, entity.ErrAlreadyExists):
			c.Status(http.StatusConflict)
		case errors.Is(err, entity.ErrNotAvailable):
			c.Status(http.StatusGone)
		default:
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusOK)
}
