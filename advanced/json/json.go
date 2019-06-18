package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	student := &Student{"张三", 20, "20190010001"}
	js, _ := json.Marshal(student)
	fmt.Printf("Json is: %s\n", js)

	bytes, err := ioutil.ReadFile("advanced/json/student.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var stu Student
	json.Unmarshal(bytes, &stu)
	fmt.Printf("stu=%v", stu)

}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	StuNo string `json:"stuNo"`
}
