package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	go func() {
		time.Sleep(10e9)
		fmt.Println("start receive data")
		x := <- c
		fmt.Println("received:", x)
	}()

	fmt.Println("ready sending data")
	for i := 0; i < 20; i++ {
		c <- i
		fmt.Println("send data:", i)
	}
	fmt.Println("sending data success")

}
