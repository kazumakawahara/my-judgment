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

func Test_updateUserUsecase_UpdateUser(t *testing.T) {
	type fields struct {
		userRepository *mockuserrepository.MockRepository
	}
	type args struct {
		ctx context.Context
		in  *userinput.UpdateUserInput
	}
	tests := []struct {
		name        string
		prepareMock func(f *fields, mockDB sqlmock.Sqlmock) error
		args        args
		want        *useroutput.UpdateUserOutput
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

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO, true).Return(userEntity, nil)

				mockDB.ExpectBegin()

				nameVO, err := uservo.NewName("ユーザー2")
				if err != nil {
					return err
				}

				idVO, err := uservo.NewNotPersistedID(0)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByName(gomock.Any(), nameVO).Return(idVO, err)

				emailVO, err := uservo.NewEmail("support2@xxxx.co.jp")
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByEmail(gomock.Any(), emailVO).Return(idVO, err)

				f.userRepository.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(nil)

				mockDB.ExpectCommit()

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.UpdateUserInput{
					UserID: 1,
					User: userinput.UpdateUser{
						Name:     testutil.ToStringPtr("ユーザー2"),
						Gender:   testutil.ToStringPtr("00101"),
						Address:  testutil.ToStringPtr("00001"),
						Email:    testutil.ToStringPtr("support2@xxxx.co.jp"),
						Password: testutil.ToStringPtr("G9OZALSRTEPXO10E"),
					},
				},
			},
			want: &useroutput.UpdateUserOutput{
				User: useroutput.UpdateUser{
					ID: 1,
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

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO, true).Return(nil, err)

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.UpdateUserInput{
					UserID: 1,
					User: userinput.UpdateUser{
						Name:     testutil.ToStringPtr("ユーザー2"),
						Gender:   testutil.ToStringPtr("00101"),
						Address:  testutil.ToStringPtr("00001"),
						Email:    testutil.ToStringPtr("support2@xxxx.co.jp"),
						Password: testutil.ToStringPtr("G9OZALSRTEPXO10E"),
					},
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound)),
		},
		{
			name: "ユーザー名重複",
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

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO, true).Return(userEntity, nil)

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
				in: &userinput.UpdateUserInput{
					UserID: 1,
					User: userinput.UpdateUser{
						Name:     testutil.ToStringPtr("ユーザー1"),
						Gender:   testutil.ToStringPtr("00101"),
						Address:  testutil.ToStringPtr("00001"),
						Email:    testutil.ToStringPtr("support2@xxxx.co.jp"),
						Password: testutil.ToStringPtr("G9OZALSRTEPXO10E"),
					},
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNameConflict)),
		},
		{
			name: "Eメール重複",
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

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO, true).Return(userEntity, nil)

				mockDB.ExpectBegin()

				nameVO, err := uservo.NewName("ユーザー2")
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
				in: &userinput.UpdateUserInput{
					UserID: 1,
					User: userinput.UpdateUser{
						Name:     testutil.ToStringPtr("ユーザー2"),
						Gender:   testutil.ToStringPtr("00101"),
						Address:  testutil.ToStringPtr("00001"),
						Email:    testutil.ToStringPtr("support1@xxxx.co.jp"),
						Password: testutil.ToStringPtr("G9OZALSRTEPXO10E"),
					},
				},
			},
			want:    nil,
			wantErr: mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserEmailConflict)),
		},
		{
			name: "ユーザー情報更新時DBエラー",
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

				f.userRepository.EXPECT().FetchUserByID(gomock.Any(), userIDVO, true).Return(userEntity, nil)

				mockDB.ExpectBegin()

				nameVO, err := uservo.NewName("ユーザー2")
				if err != nil {
					return err
				}

				idVO, err := uservo.NewNotPersistedID(0)
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByName(gomock.Any(), nameVO).Return(idVO, err)

				emailVO, err := uservo.NewEmail("support2@xxxx.co.jp")
				if err != nil {
					return err
				}

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.MjUserNotFound))

				f.userRepository.EXPECT().FetchUserIDByEmail(gomock.Any(), emailVO).Return(idVO, err)

				err = mjerr.Wrap(nil, mjerr.WithOriginError(apperr.InternalServerError))

				f.userRepository.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(err)

				mockDB.ExpectRollback()

				return nil
			},
			args: args{
				ctx: context.Background(),
				in: &userinput.UpdateUserInput{
					UserID: 1,
					User: userinput.UpdateUser{
						Name:     testutil.ToStringPtr("ユーザー2"),
						Gender:   testutil.ToStringPtr("00101"),
						Address:  testutil.ToStringPtr("00001"),
						Email:    testutil.ToStringPtr("support2@xxxx.co.jp"),
						Password: testutil.ToStringPtr("G9OZALSRTEPXO10E"),
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

			u := NewUpdateUserUsecase(f.userRepository)

			got, err := u.UpdateUser(tt.args.ctx, tt.args.in)
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("differs: (-want +got)\n%s", diff)
			}

			testutil.AssertAPIError(t, tt.wantErr, err)

			if err = mockDB.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
