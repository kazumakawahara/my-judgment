// Code generated by MockGen. DO NOT EDIT.
// Source: generate_web_token_usecase.go

// Package mocktokenusecase is a generated GoMock package.
package mocktokenusecase

import (
	context "context"
	tokeninput "my-judgment/usecase/tokenusecase/tokeninput"
	tokenoutput "my-judgment/usecase/tokenusecase/tokenoutput"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGenerateWebTokenUsecase is a mock of GenerateWebTokenUsecase interface.
type MockGenerateWebTokenUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockGenerateWebTokenUsecaseMockRecorder
}

// MockGenerateWebTokenUsecaseMockRecorder is the mock recorder for MockGenerateWebTokenUsecase.
type MockGenerateWebTokenUsecaseMockRecorder struct {
	mock *MockGenerateWebTokenUsecase
}

// NewMockGenerateWebTokenUsecase creates a new mock instance.
func NewMockGenerateWebTokenUsecase(ctrl *gomock.Controller) *MockGenerateWebTokenUsecase {
	mock := &MockGenerateWebTokenUsecase{ctrl: ctrl}
	mock.recorder = &MockGenerateWebTokenUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenerateWebTokenUsecase) EXPECT() *MockGenerateWebTokenUsecaseMockRecorder {
	return m.recorder
}

// GenerateWebToken mocks base method.
func (m *MockGenerateWebTokenUsecase) GenerateWebToken(ctx context.Context, in *tokeninput.GenerateWebTokenInput) (*tokenoutput.GenerateWebTokenOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWebToken", ctx, in)
	ret0, _ := ret[0].(*tokenoutput.GenerateWebTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateWebToken indicates an expected call of GenerateWebToken.
func (mr *MockGenerateWebTokenUsecaseMockRecorder) GenerateWebToken(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWebToken", reflect.TypeOf((*MockGenerateWebTokenUsecase)(nil).GenerateWebToken), ctx, in)
}
