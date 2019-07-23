package main

import (
	"fmt"
	"java-2-go-examples/orm/model"
	"java-2-go-examples/orm/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := mysql.GetConn()
	db.AutoMigrate(&model.Student{})
	var student model.Student
	db.Where("stu_no=?", "20190010001").First(&student)
	fmt.Println("student =", student)

	var students []model.Student
	db.Find(&students)
	fmt.Println("students =", students)

}
