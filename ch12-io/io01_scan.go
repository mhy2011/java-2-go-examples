package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("请输入姓名和年龄:")
	fmt.Scanln(&name, &age)
	fmt.Printf("name=%s, age=%d\n", name, age)
}
