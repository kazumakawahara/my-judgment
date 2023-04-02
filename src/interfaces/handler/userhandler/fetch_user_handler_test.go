package userhandler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/common/testutil"
	"my-judgment/mock/mockhandler/mockcontext"
	"my-judgment/mock/mockusecase/mockuserusecase"
	"my-judgment/usecase/userusecase/userinput"
	"my-judgment/usecase/userusecase/useroutput"
)

func Test_fetchUserHandler_FetchUser(t *testing.T) {
	type fields struct {
		fetchUserUsecase *mockuserusecase.MockFetchUserUsecase
	}
	type args struct {
		ctx *mockcontext.MockContext
	}
	tests := []struct {
		name               string
		fileSuffix         string
		args               args
		prepareMockUsecase func(f *fields)
		prepareMockRequest func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder)
	}{
		{
			name:       "正常",
			fileSuffix: "200",
			args: args{
				ctx: &mockcontext.MockContext{},
			},
			prepareMockUsecase: func(f *fields) {
				ctx := context.Background()
				in := &userinput.FetchUserInput{
					UserID: 1,
				}

				out := &useroutput.FetchUserOutput{
					User: useroutput.FetchUser{
						Name:      "ユーザー1",
						Birthday:  time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
						Gender:    "00101",
						Address:   "00001",
						Email:     "support1@xxxx.co.jp",
						Plan:      0,
						CreatedAt: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
					},
				}

				f.fetchUserUsecase.EXPECT().FetchUser(ctx, in).Return(out, nil)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				e := echo.New()
				echoCtx := e.NewContext(r, w)
				echoCtx.SetParamNames("userID")
				echoCtx.SetParamValues("1")
				ctx.Context = echoCtx
			},
		},
		{
			name:       "ユーザーID不正",
			fileSuffix: "400",
			args: args{
				ctx: &mockcontext.MockContext{},
			},
			prepareMockUsecase: nil,
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				e := echo.New()
				echoCtx := e.NewContext(r, w)
				echoCtx.SetParamNames("userID")
				echoCtx.SetParamValues("InvalidParameter request")
				ctx.Context = echoCtx
			},
		},
		{
			name:       "usecase層での404エラー(指定されたユーザーが存在しない場合)",
			fileSuffix: "404",
			args: args{
				ctx: &mockcontext.MockContext{},
			},
			prepareMockUsecase: func(f *fields) {
				ctx := context.Background()
				in := &userinput.FetchUserInput{
					UserID: 1,
				}

				err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.fetchUserUsecase.EXPECT().FetchUser(ctx, in).Return(nil, err)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				e := echo.New()
				echoCtx := e.NewContext(r, w)
				echoCtx.SetParamNames("userID")
				echoCtx.SetParamValues("1")
				ctx.Context = echoCtx
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				fetchUserUsecase: mockuserusecase.NewMockFetchUserUsecase(gmctrl),
			}

			if tt.prepareMockUsecase != nil {
				tt.prepareMockUsecase(&f)
			}

			h := NewFetchUserHandler(f.fetchUserUsecase)

			r := httptest.NewRequest(http.MethodGet, "/mj/users/:userID", nil)
			w := httptest.NewRecorder()

			if tt.prepareMockRequest != nil {
				tt.prepareMockRequest(tt.args.ctx, r, w)
			}

			if err := h.FetchUser(tt.args.ctx); err != nil {
				t.Fatalf("FetchUser() error = %v", err)
			}

			res := w.Result()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
