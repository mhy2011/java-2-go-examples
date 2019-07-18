package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	if rs, err := delive(20, -5); err == nil{
		fmt.Printf("%f / %f = %f\n", 20, -5, rs)
	} else {
		fmt.Println(err)
	}
	if rs, err := delive(20, 0); err == nil{
		fmt.Printf("%f / %f = %f\n", 20, -5, rs)
	} else {
		fmt.Println(err)
	}
}

func delive(a, b float32) (float32, error) {
	if b < 0 {
		return 0, errors.New("被除数不能为负数")
	} else if b == 0 {
		panic("被除能为0")
	}
	return a / b, nil
}
