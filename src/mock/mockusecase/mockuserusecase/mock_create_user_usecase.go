// Code generated by MockGen. DO NOT EDIT.
// Source: create_user_usecase.go

// Package mockuserusecase is a generated GoMock package.
package mockuserusecase

import (
	context "context"
	userinput "my-judgment/usecase/userusecase/userinput"
	useroutput "my-judgment/usecase/userusecase/useroutput"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCreateUserUsecase is a mock of CreateUserUsecase interface.
type MockCreateUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateUserUsecaseMockRecorder
}

// MockCreateUserUsecaseMockRecorder is the mock recorder for MockCreateUserUsecase.
type MockCreateUserUsecaseMockRecorder struct {
	mock *MockCreateUserUsecase
}

// NewMockCreateUserUsecase creates a new mock instance.
func NewMockCreateUserUsecase(ctrl *gomock.Controller) *MockCreateUserUsecase {
	mock := &MockCreateUserUsecase{ctrl: ctrl}
	mock.recorder = &MockCreateUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateUserUsecase) EXPECT() *MockCreateUserUsecaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockCreateUserUsecase) CreateUser(ctx context.Context, in *userinput.CreateUserInput) (*useroutput.CreateUserOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, in)
	ret0, _ := ret[0].(*useroutput.CreateUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockCreateUserUsecaseMockRecorder) CreateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockCreateUserUsecase)(nil).CreateUser), ctx, in)
}
