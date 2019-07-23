package main

import (
	_ "github.com/go-sql-driver/mysql"
	"java-2-go-examples/orm/model"
	"java-2-go-examples/orm/mysql"
	"time"
)

func main() {
	db := mysql.GetConn()

	//开启事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()	//事务回滚
		}
	}()

	student := model.Student{Name: "李四", StuNo: "20190020002", Gender: "M", RegDate: time.Now()}
	studentScore := model.StudentScore{StuNo: "20190020002", Course: "数学", Score: 100}

	if e := tx.Create(&student).Error; e != nil {
		tx.Rollback()	//回滚事务
	}
	if e := tx.Create(&studentScore).Error; e != nil {
		tx.Rollback() //回滚事务
	}
	tx.Commit()	//提交事务
}
