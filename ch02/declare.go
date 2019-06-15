package main

import "fmt"

// 声明常量
const PI float64 = 3.14

// 声明函数
func main() {
	// 声明变量
	var radius float64 = 10
	area := PI * radius * radius
	fmt.Println("area is :", area)
}
