package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //读取键盘输入
	fmt.Println("请输入内容:")
	s, e := reader.ReadString('\n') //读取输入，直到换行
	if e != nil {
		fmt.Printf("read fail. e=%s\n", e)
		return
	}

	fmt.Println("输入的内容为:", s)
}
