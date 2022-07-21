package databases

import (
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

func InitMysql(dsn string) (err error) {
	db, err := gorm.Open("msyql", dsn)
	if err != nil {
		return err
	}
	DB = db
	return
}
