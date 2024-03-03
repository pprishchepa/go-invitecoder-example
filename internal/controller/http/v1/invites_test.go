package v1_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	httpctrl "github.com/pprishchepa/go-invitecoder-example/internal/controller/http"
	v1 "github.com/pprishchepa/go-invitecoder-example/internal/controller/http/v1"
	"github.com/pprishchepa/go-invitecoder-example/internal/entity"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestInvitesRoutes(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	type args struct {
		data             string
		mockExpectations func(svc *MockInviteService)
	}
	type want struct {
		status int
	}
	cases := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid request",
			args: args{
				data: `{"email": "foo@example.com", "code": "twitter"}`,
				mockExpectations: func(svc *MockInviteService) {
					svc.EXPECT().
						AcceptInvite(gomock.Any(), entity.InvitedUser{Email: "foo@example.com", InvitedVia: "twitter"}).
						Return(nil)
				},
			},
			want: want{
				status: http.StatusOK,
			},
		},
		{
			name: "invalid request - already exists",
			args: args{
				data: `{"email": "foo@example.com", "code": "twitter"}`,
				mockExpectations: func(svc *MockInviteService) {
					svc.EXPECT().
						AcceptInvite(gomock.Any(), gomock.Any()).
						Return(entity.ErrAlreadyExists)
				},
			},
			want: want{
				status: http.StatusConflict,
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			data := bytes.NewBufferString(tt.args.data)
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/v1/invitation", data)
			require.Empty(t, err)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			service := NewMockInviteService(mockCtrl)
			tt.args.mockExpectations(service)

			routes := v1.NewInvitesRoutes(service, zerolog.Nop())
			router := httpctrl.NewRouter(routes)

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if tt.want.status != 0 {
				require.Equal(t, tt.want.status, rr.Code)
			}
		})
	}
}
