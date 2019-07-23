package main

import (
	_ "github.com/go-sql-driver/mysql"
	"java-2-go-examples/orm/model"
	"java-2-go-examples/orm/mysql"
)

func main() {
	db := mysql.GetConn()
	db.Model(&model.Student{}).Where("stu_no=?", "20190010001").
		Update("gender", "W")
}
