package main

import (
	"fmt"
	"time"
)

func main() {
	strCh := make(chan string)
	go producer01(strCh)
	fmt.Println(<-strCh)
	time.Sleep(5e9) //休眠5s
}

func producer01(ch chan string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("Hello_%d", i)
	}
}
