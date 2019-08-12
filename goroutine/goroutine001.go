package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start execute main func")
	go printA() //启动协程
	go printB() //启动协程
	time.Sleep(1 * time.Second)
	fmt.Println("finish main func")
}

func printA() {
	for i := 0; i < 5; i++ {
		fmt.Printf("a_%d\n", i)
		time.Sleep(10 * time.Millisecond) //休眠10毫秒
	}
}
func printB() {
	for i := 0; i < 5; i++ {
		fmt.Printf("b_%d\n", i)
		time.Sleep(10 * time.Millisecond) //休眠10毫秒
	}
}
