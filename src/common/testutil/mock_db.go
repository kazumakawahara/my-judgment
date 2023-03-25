package testutil

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, nil, err
	}

	db, err := gorm.Open(
		mysql.Dialector{
			Config: &mysql.Config{
				DriverName:                "mysql",
				Conn:                      mockDB,
				SkipInitializeWithVersion: true,
			},
		},
	)
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}

func CloseMockDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
