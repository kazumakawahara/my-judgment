package userusecase

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"my-judgment/common/apperr"
	"my-judgment/common/config"
	"my-judgment/common/mjerr"
	"my-judgment/common/testutil"
	"my-judgment/common/vo/uservo"
	"my-judgment/mock/mockrepository/mockuserrepository"
	"my-judgment/usecase/userusecase/userinput"
	"my-judgment/usecase/userusecase/useroutput"
)

func Test_createUserUsecase_CreateUser(t *testing.T) {
	type fields struct {
		userRepository *mockuserrepository.MockRepository
	}
	type args struct {
		ctx context.Context
		in  *userinput.CreateUserInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields, mockDB sqlmock.Sqlmock) error
		args        args
		want        *useroutput.CreateUserOutput
		wantErr     error
	}{
		{
			name: "正常",
			prepareMock: func(f *fields, mockDB sqlmock.Sqlmock) error {
				f.userRepository.EXPECT().ExistsUserByPassword(gomock.Any(), gomock.Any()).Return(false, nil)

				mockDB.ExpectBegin()

				nameVO, err := uservo.NewName("ユーザー1")
				if err != nil {
					return err
				}

				idVO, err := uservo.NewNotPersistedID(0)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByName(gomock.Any(), nameVO).Return(idVO, err)

				emailVO, err := uservo.NewEmail("support1@xxxx.co.jp")
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByEmail(gomock.Any(), emailVO).Return(idVO, err)

				idVO, err = uservo.NewID(1)
				if err != nil {
					return err
				}

				f.userRepository.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(idVO, nil)

				mockDB.ExpectCommit()

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.CreateUserInput{
					User: userinput.CreateUser{
						Name:     "ユーザー1",
						Birthday: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
						Gender:   "00101",
						Address:  "00001",
						Email:    "support1@xxxx.co.jp",
					},
				},
			},
			want: &useroutput.CreateUserOutput{
				User: useroutput.CreateUser{
					ID:       1,
					Password: "G9OZALSRTEPXO10E",
				},
			},
			wantErr: nil,
		},
		{
			name: "password重複確認時DBエラー",
			prepareMock: func(f *fields, mockDB sqlmock.Sqlmock) error {
				err := mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError))

				f.userRepository.EXPECT().ExistsUserByPassword(gomock.Any(), gomock.Any()).Return(false, err)

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.CreateUserInput{
					User: userinput.CreateUser{
						Name:     "ユーザー1",
						Birthday: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
						Gender:   "00101",
						Address:  "00001",
						Email:    "support1@xxxx.co.jp",
					},
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError)),
		},
		{
			name: "ユーザー名重複",
			prepareMock: func(f *fields, mockDB sqlmock.Sqlmock) error {
				f.userRepository.EXPECT().ExistsUserByPassword(gomock.Any(), gomock.Any()).Return(false, nil)

				mockDB.ExpectBegin()

				nameVO, err := uservo.NewName("ユーザー1")
				if err != nil {
					return err
				}

				idVO, err := uservo.NewID(1)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNameConflict))

				f.userRepository.EXPECT().FetchUserIDByName(gomock.Any(), nameVO).Return(idVO, err)

				mockDB.ExpectRollback()

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.CreateUserInput{
					User: userinput.CreateUser{
						Name:     "ユーザー1",
						Birthday: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
						Gender:   "00101",
						Address:  "00001",
						Email:    "support1@xxxx.co.jp",
					},
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNameConflict)),
		},
		{
			name: "Eメール重複",
			prepareMock: func(f *fields, mockDB sqlmock.Sqlmock) error {
				f.userRepository.EXPECT().ExistsUserByPassword(gomock.Any(), gomock.Any()).Return(false, nil)

				mockDB.ExpectBegin()

				nameVO, err := uservo.NewName("ユーザー1")
				if err != nil {
					return err
				}

				idVO, err := uservo.NewNotPersistedID(0)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByName(gomock.Any(), nameVO).Return(idVO, err)

				emailVO, err := uservo.NewEmail("support1@xxxx.co.jp")
				if err != nil {
					return err
				}

				idVO, err = uservo.NewID(1)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserEmailConflict))

				f.userRepository.EXPECT().FetchUserIDByEmail(gomock.Any(), emailVO).Return(idVO, err)

				mockDB.ExpectRollback()

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.CreateUserInput{
					User: userinput.CreateUser{
						Name:     "ユーザー1",
						Birthday: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
						Gender:   "00101",
						Address:  "00001",
						Email:    "support1@xxxx.co.jp",
					},
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserEmailConflict)),
		},
		{
			name: "ユーザー新規登録時DBエラー",
			prepareMock: func(f *fields, mockDB sqlmock.Sqlmock) error {
				f.userRepository.EXPECT().ExistsUserByPassword(gomock.Any(), gomock.Any()).Return(false, nil)

				mockDB.ExpectBegin()

				nameVO, err := uservo.NewName("ユーザー1")
				if err != nil {
					return err
				}

				idVO, err := uservo.NewNotPersistedID(0)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByName(gomock.Any(), nameVO).Return(idVO, err)

				emailVO, err := uservo.NewEmail("support1@xxxx.co.jp")
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByEmail(gomock.Any(), emailVO).Return(idVO, err)

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError))

				f.userRepository.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(idVO, err)

				mockDB.ExpectRollback()

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.CreateUserInput{
					User: userinput.CreateUser{
						Name:     "ユーザー1",
						Birthday: time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC),
						Gender:   "00101",
						Address:  "00001",
						Email:    "support1@xxxx.co.jp",
					},
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

			u := NewCreateUserUsecase(f.userRepository)

			got, err := u.CreateUser(tt.args.ctx, tt.args.in)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(useroutput.CreateUser{}, "Password")); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			testutil.AssertAPIError(t, tt.wantErr, err)

			if err = mockDB.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
