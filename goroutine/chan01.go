package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start main")
	strCh := make(chan string)
	go producer(strCh)  //启动生产者
	go consumer(strCh)  //启动消费者
	time.Sleep(6 * time.Second) //休眠6s
	fmt.Println("end main")
}

func producer(ch chan string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("ready to produce msg, i  = %d\n", i)
		ch <- fmt.Sprintf("hello_%d", i)
	}
}

func consumer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("consume data :", data)
	}
}
