package main

import (
	"fmt"
	"time"
)

func main() {
	go printMsg(500 * time.Millisecond) //启动一个新线程
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func printMsg(delay time.Duration) {
	fmt.Printf("程序执行中")
	for {
		fmt.Printf(".")
		time.Sleep(delay) //休眠一下
	}
}
