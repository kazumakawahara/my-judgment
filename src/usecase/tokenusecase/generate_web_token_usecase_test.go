package tokenusecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
	"my-judgment/common/testutil"
	"my-judgment/common/vo/uservo"
	"my-judgment/domain/userdm"
	"my-judgment/mock/mockrepository/mockuserrepository"
	"my-judgment/mock/mocktokenservice/mocktokenservice"
	"my-judgment/usecase/tokenusecase/tokeninput"
	"my-judgment/usecase/tokenusecase/tokenoutput"
)

func Test_generateWebTokenUsecase_GenerateWebToken(t *testing.T) {
	type fields struct {
		tokenService   *mocktokenservice.MockTokenService
		userRepository *mockuserrepository.MockRepository
	}
	type args struct {
		ctx context.Context
		in  *tokeninput.GenerateWebTokenInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields) error
		args        args
		want        *tokenoutput.GenerateWebTokenOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields) error {
				ClientToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A" //#nosec G101 - This is a false positive

				userID := 100

				f.tokenService.EXPECT().ParseWebClientToken(ClientToken).Return(userID, nil)

				userIDVO, err := uservo.NewID(userID)
				if err != nil {
					return err
				}

				userEntity, err := userdm.Reconstruct(
					1,
					"ユーザー1",
					time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
					"00101",
					"00001",
					"support1@xxxx.co.jp",
					"G9OZALSRTEPXO10E",
					0,
					time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
					100,
					time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
					100,
					nil,
					nil,
				)
				if err != nil {
					return err
				}

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO, false).Return(userEntity, nil)

				token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteUp1ZGdtZW50IiwiYXVkIjoiTVRBd01BPT0iLCJ1c2VySUQiOjEwMDAsImV4cCI6MTY4NTg3Nzk0NCwiaWF0IjoxNjg1ODc0MzQ0fQ.ZbxHQGTK-zgfl2wNbNFhWS2pjZcxKGkwBhUsPNwZwt8" //#nosec G101 - This is a false positive

				f.tokenService.EXPECT().GenerateWebAuthToken(userIDVO, gomock.Any()).Return(token, nil)

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &tokeninput.GenerateWebTokenInput{
					ClientToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A",
				},
			},
			want: &tokenoutput.GenerateWebTokenOutput{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJteUp1ZGdtZW50IiwiYXVkIjoiTVRBd01BPT0iLCJ1c2VySUQiOjEwMDAsImV4cCI6MTY4NTg3Nzk0NCwiaWF0IjoxNjg1ODc0MzQ0fQ.ZbxHQGTK-zgfl2wNbNFhWS2pjZcxKGkwBhUsPNwZwt8",
			},
			wantErr: nil,
		},
		{
			name: "Authorization header不正",
			prepareMock: func(f *fields) error {
				ClientToken := "Invalid request"

				err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InvalidRequestToken))

				f.tokenService.EXPECT().ParseWebClientToken(ClientToken).Return(0, err)

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &tokeninput.GenerateWebTokenInput{
					ClientToken: "Invalid request",
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InvalidRequestToken)),
		},
		{
			name: "指定されたユーザーが存在しない場合",
			prepareMock: func(f *fields) error {
				ClientToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A" //#nosec G101 - This is a false positive

				userID := 100

				f.tokenService.EXPECT().ParseWebClientToken(ClientToken).Return(userID, nil)

				userIDVO, err := uservo.NewID(userID)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO, false).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &tokeninput.GenerateWebTokenInput{
					ClientToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A",
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound)),
		},
		{
			name: "token発行時DBエラー",
			prepareMock: func(f *fields) error {
				ClientToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A" //#nosec G101 - This is a false positive

				userID := 100

				f.tokenService.EXPECT().ParseWebClientToken(ClientToken).Return(userID, nil)

				userIDVO, err := uservo.NewID(userID)
				if err != nil {
					return err
				}

				userEntity, err := userdm.Reconstruct(
					1,
					"ユーザー1",
					time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
					"00101",
					"00001",
					"support1@xxxx.co.jp",
					"G9OZALSRTEPXO10E",
					0,
					time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
					100,
					time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
					100,
					nil,
					nil,
				)
				if err != nil {
					return err
				}

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO, false).Return(userEntity, nil)

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError))

				f.tokenService.EXPECT().GenerateWebAuthToken(userIDVO, gomock.Any()).Return("", err)

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &tokeninput.GenerateWebTokenInput{
					ClientToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJNVEF3TUE9PSIsImlzcyI6Im15SnVkZ21lbnQifQ.5suDqrpLNU7YnZohWJG3i5DLlTxpXbq8zimSU-T0K6A",
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				tokenService:   mocktokenservice.NewMockTokenService(gmctrl),
				userRepository: mockuserrepository.NewMockRepository(gmctrl),
			}

			if tt.prepareMock != nil {
				if err := tt.prepareMock(&f); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			u := NewGenerateWebTokenUsecase(f.tokenService, f.userRepository)

			got, err := u.GenerateWebToken(tt.args.ctx, tt.args.in)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("GenerateWebToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			testutil.AssertAPIError(t, tt.wantErr, err)
		})
	}
}
