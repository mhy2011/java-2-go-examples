package main

import "fmt"

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度

const AbsoluteZeroC Celsius = -273.15 //绝对零度
const FreezingZeroC Celsius = 0       //结冰温度
const BoilingC Celsius = 100          //沸水温度

func main() {
	fa := ce2fa(AbsoluteZeroC)
	fmt.Println("fa is :", fa)
}

func ce2fa(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func fa2ce(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
