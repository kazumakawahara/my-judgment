package tokenhandler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/common/testutil"
	"my-judgment/mock/mockhandler/mockcontext"
	"my-judgment/mock/mockusecase/mocktokenusecase"
	"my-judgment/usecase/tokenusecase/tokeninput"
	"my-judgment/usecase/tokenusecase/tokenoutput"
)

func Test_generateWebTokenHandler_GenerateWebToken(t *testing.T) {
	type fields struct {
		generateTokenUsecase *mocktokenusecase.MockGenerateWebTokenUsecase
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
				in := &tokeninput.GenerateWebTokenInput{
					ClientToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A",
				}

				out := &tokenoutput.GenerateWebTokenOutput{
					Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteUp1ZGdtZW50IiwiYXVkIjoiTVRBd01BPT0iLCJ1c2VySUQiOjEwMDAsImV4cCI6MTY4NTg3Nzk0NCwiaWF0IjoxNjg1ODc0MzQ0fQ.ZbxHQGTK-zgfl2wNbNFhWS2pjZcxKGkwBhUsPNwZwt8",
				}

				f.generateTokenUsecase.EXPECT().GenerateWebToken(ctx, in).Return(out, nil)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				r.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A")

				e := echo.New()
				echoCtx := e.NewContext(r, w)
				ctx.Context = echoCtx
			},
		},
		{
			name:       "Authorization header不正",
			fileSuffix: "400",
			args: args{
				ctx: &mockcontext.MockContext{},
			},
			prepareMockUsecase: func(f *fields) {
				ctx := context.Background()
				in := &tokeninput.GenerateWebTokenInput{
					ClientToken: "Invalid request",
				}

				err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InvalidRequestToken))

				f.generateTokenUsecase.EXPECT().GenerateWebToken(ctx, in).Return(nil, err)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				r.Header.Set("Authorization", "Invalid request")

				e := echo.New()
				echoCtx := e.NewContext(r, w)
				ctx.Context = echoCtx
			},
		},
		{
			name:       "usecase層での404エラー(ユーザーが存在しない場合)",
			fileSuffix: "404",
			args: args{
				ctx: &mockcontext.MockContext{},
			},
			prepareMockUsecase: func(f *fields) {
				ctx := context.Background()
				in := &tokeninput.GenerateWebTokenInput{
					ClientToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A",
				}

				err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.generateTokenUsecase.EXPECT().GenerateWebToken(ctx, in).Return(nil, err)
			},
			prepareMockRequest: func(ctx *mockcontext.MockContext, r *http.Request, w *httptest.ResponseRecorder) {
				r.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A")

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
				generateTokenUsecase: mocktokenusecase.NewMockGenerateWebTokenUsecase(gmctrl),
			}

			if tt.prepareMockUsecase != nil {
				tt.prepareMockUsecase(&f)
			}

			h := NewGenerateWebTokenHandler(f.generateTokenUsecase)

			r := httptest.NewRequest(http.MethodGet, "/mj/token", nil)
			w := httptest.NewRecorder()

			if tt.prepareMockRequest != nil {
				tt.prepareMockRequest(tt.args.ctx, r, w)
			}

			if err := h.GenerateWebToken(tt.args.ctx); err != nil {
				t.Fatalf("GenerateWebToken() error = %v", err)
			}

			res := w.Result()

			testutil.AssertResponseBody(t, res, tt.fileSuffix)
		})
	}
}
