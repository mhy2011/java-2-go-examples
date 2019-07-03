package main

import (
	"fmt"
	"time"
)

func main() {
	strCh := make(chan string)
	go producer02(strCh)
	go consumer02(strCh)
	time.Sleep(5e9) //休眠5s
}

func producer02(ch chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(2e8)
		ch <- fmt.Sprintf("Hello_%d", i)
	}
}

func consumer02(ch chan string) {
	for ; ; {
		fmt.Println(<-ch)
	}
}
