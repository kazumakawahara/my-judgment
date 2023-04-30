package userhandler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

func Test_updateUserHandler_UpdateUser(t *testing.T) {
	type fields struct {
		updateUserUsecase *mockuserusecase.MockUpdateUserUsecase
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
				in := &userinput.UpdateUserInput{
					UserID: 1,
					User: userinput.UpdateUser{
						Name:     testutil.ToStringPtr("ユーザー2"),
						Gender:   testutil.ToStringPtr("00101"),
						Address:  testutil.ToStringPtr("00001"),
						Email:    testutil.ToStringPtr("support1@xxxx.co.jp"),
						Password: testutil.ToStringPtr("G9OZALSRTEPXO10E"),
					},
				}

				out := &useroutput.UpdateUserOutput{
					User: useroutput.UpdateUser{
						ID: 1,
					},
				}

				f.updateUserUsecase.EXPECT().UpdateUser(ctx, in).Return(out, nil)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				r.Header.Set("Content-Type", "application/json")

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
				r.Header.Set("Content-Type", "application/json")

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
				in := &userinput.UpdateUserInput{
					UserID: 1,
					User: userinput.UpdateUser{
						Name:     testutil.ToStringPtr("ユーザー2"),
						Gender:   testutil.ToStringPtr("00101"),
						Address:  testutil.ToStringPtr("00001"),
						Email:    testutil.ToStringPtr("support1@xxxx.co.jp"),
						Password: testutil.ToStringPtr("G9OZALSRTEPXO10E"),
					},
				}

				err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.updateUserUsecase.EXPECT().UpdateUser(ctx, in).Return(nil, err)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				r.Header.Set("Content-Type", "application/json")

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
				updateUserUsecase: mockuserusecase.NewMockUpdateUserUsecase(gmctrl),
			}

			if tt.prepareMockUsecase != nil {
				tt.prepareMockUsecase(&f)
			}

			h := NewUpdateUserHandler(f.updateUserUsecase)

			r := httptest.NewRequest(http.MethodPut, "/mj/users/:userID", strings.NewReader(testutil.GetRequestJsonFromTestData(t, tt.fileSuffix)))
			w := httptest.NewRecorder()

			if tt.prepareMockRequest != nil {
				tt.prepareMockRequest(tt.args.ctx, r, w)
			}

			if err := h.UpdateUser(tt.args.ctx); err != nil {
				t.Fatalf("UpdateUser() error = %v", err)
			}

			res := w.Result()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
