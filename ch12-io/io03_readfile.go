package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, e := os.Open("ch12-io/io03_readfile.go")
	if e != nil {
		fmt.Println("open file error.", e)
		return
	}
	defer file.Close() //关闭文件

	reader := bufio.NewReader(file)
	for {
		s, e := reader.ReadString('\n')
		fmt.Printf(s)
		if e == io.EOF {
			return
		}
	}
}
