package databases

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

func InitMysql(dsn string) (err error) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	Db = db
	return
}
