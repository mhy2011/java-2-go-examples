package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"java-2-go-examples/advanced/conf"
)

var (
	env    string //环境信息
	Config conf.Config
)

func init() {
	// 通过命令行传递环境信息
	flag.StringVar(&env, "env", "dev", "which env you want")
	flag.Parse()
	// 初始化配置信息
	initConfig(env)
}

func main() {
	fmt.Println("start application.")
	fmt.Println("config is:", Config)
}

// 初始化配置信息
func initConfig(env string) error {
	bytes, err := ioutil.ReadFile("advanced/conf/" + env + "/config.json")
	if err != nil {
		fmt.Println("read file error .", err)
		return err
	}
	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		fmt.Println("parse json error .", err)
		return err
	}
	return nil
}
