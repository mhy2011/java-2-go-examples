package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})
	http.HandleFunc("/hello", hello)	//设置访问路由
	err := http.ListenAndServe(":8080", nil)	//设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello")	//输出到客户端
}