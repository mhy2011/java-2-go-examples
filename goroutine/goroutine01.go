package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main func start")
	go wait() //启动协程
	fmt.Println("ready to sleep for main")
	time.Sleep(6 * 1e9) //主协程休眠6s
	fmt.Println("main func end")
}

func wait() {
	fmt.Println("wait func start")
	time.Sleep(3 * 1e9) //休眠3s
	fmt.Println("wait func end")
}
