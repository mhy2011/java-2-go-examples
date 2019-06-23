package main

import "fmt"

func main() {
	r := Rectangle{10, 5}
	shapers := []Shaper{r}
	for _, val := range shapers {
		fmt.Printf("shape is %v, area is %f\n", val, val.Area())
	}
}

// 定义接口
type Shaper interface {
	Area() float32
}

type Rectangle struct {
	length float32 //长
	width  float32 //宽
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}
