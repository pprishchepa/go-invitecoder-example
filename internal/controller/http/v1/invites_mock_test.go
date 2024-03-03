// Code generated by MockGen. DO NOT EDIT.
// Source: invites.go
//
// Generated by this command:
//
//	mockgen -source=invites.go -destination=invites_mock_test.go -package=v1_test
//

// Package v1_test is a generated GoMock package.
package v1_test

import (
	context "context"
	reflect "reflect"

	entity "github.com/pprishchepa/go-invitecoder-example/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockInviteService is a mock of InviteService interface.
type MockInviteService struct {
	ctrl     *gomock.Controller
	recorder *MockInviteServiceMockRecorder
}

// MockInviteServiceMockRecorder is the mock recorder for MockInviteService.
type MockInviteServiceMockRecorder struct {
	mock *MockInviteService
}

// NewMockInviteService creates a new mock instance.
func NewMockInviteService(ctrl *gomock.Controller) *MockInviteService {
	mock := &MockInviteService{ctrl: ctrl}
	mock.recorder = &MockInviteServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInviteService) EXPECT() *MockInviteServiceMockRecorder {
	return m.recorder
}

// AcceptInvite mocks base method.
func (m *MockInviteService) AcceptInvite(ctx context.Context, user entity.InvitedUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptInvite", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptInvite indicates an expected call of AcceptInvite.
func (mr *MockInviteServiceMockRecorder) AcceptInvite(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptInvite", reflect.TypeOf((*MockInviteService)(nil).AcceptInvite), ctx, user)
}
