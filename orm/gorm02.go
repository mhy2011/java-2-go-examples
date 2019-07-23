package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"java-2-go-examples/orm/model"
	"time"
)

var db *gorm.DB

func init()  {
	mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", "root", "root", "tcp", "127.0.0.1", 3306, "go-examples")
	db, _ = gorm.Open("mysql", mysql)
	db.LogMode(true)
	//defer db.Close()
}

func main() {
	db.AutoMigrate(&model.Student{})
	// 创建
	db.Create(&model.Student{StuNo:"20190010001", Name:"张三", Gender:"M", RegDate: time.Now()})
	// 查询
	var student model.Student
	db.First(&student, 1)
	fmt.Println("student =", student)
}
