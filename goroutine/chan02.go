package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start main")
	chanSize := 5
	//带缓冲区的通道
	strCh := make(chan string, chanSize)

	go producerWithBuffer(strCh)
	time.Sleep(3 * time.Second) //主线程休眠3s
	fmt.Println("end main")
}

func producerWithBuffer(strCh chan string) {
	for i := 0; i < 20; i++ {
		strCh <- fmt.Sprintf("hello_%d", i)
		fmt.Println("send message success i =", i)
	}
}