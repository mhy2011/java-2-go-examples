package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

func init()  {
	mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", "root", "root", "tcp", "127.0.0.1", 3306, "go-examples")
	db, _ = gorm.Open("mysql", mysql)
	db.LogMode(true)
	//defer db.Close()
}


func main() {
	db.AutoMigrate(&Product{})
	//db.CreateTable(&Product{})
	// 创建
	//db.Create(&Product{Code:"P123456", Price:200})
	//读取
	var product Product
	db.First(&product, 1)	//查询Id为1的Product
	fmt.Println("product=", product)
	db.First(&product, "code=?", "P123456")
	fmt.Println("product=", product)

	// 更新
	db.Model(&product).Update("Price", 300)

	//删除
	db.Delete(&product)
}

type Product struct {
	gorm.Model
	Code string
	Price uint
}

type User struct {
	gorm.Model
	Name string
	Age uint
	Birthday *time.Time
	Email string `gorm:unique_index`

}