package userusecase

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"my-judgment/common/apperr"
	"my-judgment/common/config"
	"my-judgment/common/mjerr"
	"my-judgment/common/testutil"
	"my-judgment/common/vo/uservo"
	"my-judgment/domain/userdm"
	"my-judgment/mock/mockrepository/mockuserrepository"
	"my-judgment/usecase/userusecase/userinput"
	"my-judgment/usecase/userusecase/useroutput"
)

func Test_fetchUserUsecase_FetchUser(t *testing.T) {
	type fields struct {
		userRepository *mockuserrepository.MockRepository
	}
	type args struct {
		ctx context.Context
		in  *userinput.FetchUserInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields, mockDB sqlmock.Sqlmock) error
		args        args
		want        *useroutput.FetchUserOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields, mockDB sqlmock.Sqlmock) error {
				userIDVO, err := uservo.NewID(1)
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

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO).Return(userEntity, nil)

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.FetchUserInput{
					UserID: 1,
				},
			},
			want: &useroutput.FetchUserOutput{
				User: useroutput.FetchUser{
					Name:      "ユーザー1",
					Birthday:  time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
					Gender:    "00101",
					Address:   "00001",
					Email:     "support1@xxxx.co.jp",
					Plan:      0,
					CreatedAt: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "指定されたユーザーが存在しない場合",
			prepareMock: func(f *fields, mockDB sqlmock.Sqlmock) error {
				userIDVO, err := uservo.NewID(1)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.FetchUserInput{
					UserID: 1,
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmctrl := gomock.NewController(t)

			f := fields{
				userRepository: mockuserrepository.NewMockRepository(gmctrl),
			}

			db, mockDB, err := testutil.NewMockDB()
			if err != nil {
				t.Fatalf("error '%s' was not expected when opening a stub database connection", err)
			}
			defer testutil.CloseMockDB(db)

			if tt.prepareMock != nil {
				if err = tt.prepareMock(&f, mockDB); err != nil {
					t.Fatalf("prepareMock() error = %v", err)
				}
			}

			tt.args.ctx = context.WithValue(tt.args.ctx, config.DBKey, db)

			u := NewFetchUserUsecase(f.userRepository)

			got, err := u.FetchUser(tt.args.ctx, tt.args.in)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("FetchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			testutil.AssertAPIError(t, tt.wantErr, err)
		})
	}
}
