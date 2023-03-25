package userhandler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
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

func Test_createUserHandler_CreateUser(t *testing.T) {
	type fields struct {
		createUserUsecase *mockuserusecase.MockCreateUserUsecase
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
				in := &userinput.CreateUserInput{
					User: userinput.CreateUser{
						Name:     "ユーザー1",
						Birthday: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
						Gender:   "00101",
						Address:  "00001",
						Email:    "support1@xxxx.co.jp",
					},
				}

				out := &useroutput.CreateUserOutput{
					User: useroutput.CreateUser{
						ID:       1,
						Password: "G9OZALSRTEPXO10E",
					},
				}

				f.createUserUsecase.EXPECT().CreateUser(ctx, in).Return(out, nil)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				r.Header.Set("Content-Type", "application/json")

				e := echo.New()
				echoCtx := e.NewContext(r, w)
				ctx.Context = echoCtx
			},
		},
		{
			name:       "Bindエラー",
			fileSuffix: "400",
			args: args{
				ctx: &mockcontext.MockContext{},
			},
			prepareMockUsecase: nil,
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				r.Header.Set("Content-Type", "application/json")

				e := echo.New()
				echoCtx := e.NewContext(r, w)
				ctx.Context = echoCtx
			},
		},
		{
			name:       "usecase層での409エラー(ユーザー名重複)",
			fileSuffix: "409",
			args: args{
				ctx: &mockcontext.MockContext{},
			},
			prepareMockUsecase: func(f *fields) {
				ctx := context.Background()
				in := &userinput.CreateUserInput{
					User: userinput.CreateUser{
						Name:     "ユーザー1",
						Birthday: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
						Gender:   "00101",
						Address:  "00001",
						Email:    "support1@xxxx.co.jp",
					},
				}

				err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNameConflict))

				f.createUserUsecase.EXPECT().CreateUser(ctx, in).Return(nil, err)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				r.Header.Set("Content-Type", "application/json")

				e := echo.New()
				echoCtx := e.NewContext(r, w)
				ctx.Context = echoCtx
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				createUserUsecase: mockuserusecase.NewMockCreateUserUsecase(gmctrl),
			}

			if tt.prepareMockUsecase != nil {
				tt.prepareMockUsecase(&f)
			}

			h := NewCreateUserHandler(f.createUserUsecase)

			r := httptest.NewRequest(http.MethodPost, "/mj/users", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			if tt.prepareMockRequest != nil {
				tt.prepareMockRequest(tt.args.ctx, r, w)
			}

			if err := h.CreateUser(tt.args.ctx); err != nil {
				t.Fatalf("CreateUser() error = %v", err)
			}

			res := w.Result()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
