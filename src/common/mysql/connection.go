package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	protocol := os.Getenv("MYSQL_PROTOCOL")
	dbname := os.Getenv("MYSQL_DATABASE")

	dsn := user + ":" + pass + "@" + protocol + "/" + dbname + "?parseTime=true&charset=utf8"

	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: l,
		},
	)
	if err != nil {
		fmt.Printf("DB Open Error :%v", err)
		panic(err.Error())
	}

	return db
}
