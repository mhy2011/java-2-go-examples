package main

import (
	"flag"
	"fmt"
)

var (
	env  string //环境信息，例如dev/pre/prod等
	port int    //端口
)

func main() {
	flag.StringVar(&env, "env", "dev", "which env you want")
	flag.IntVar(&port, "port", 8080, "which port you want")
	flag.Parse()
	fmt.Println("env =", env)
	fmt.Println("port =", port)
}
