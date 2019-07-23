package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func GetConn() *gorm.DB {
	mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", "root", "root", "tcp", "127.0.0.1", 3306, "go-examples")
	db, err := gorm.Open("mysql", mysql)
	if err != nil {
		panic("connect mysql fail")
	}
	db.LogMode(true)
	return db
}
