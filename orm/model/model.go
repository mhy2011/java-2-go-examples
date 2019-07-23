package model

import "time"

type Student struct {
	ID uint64
	Name string
	StuNo string
	Gender string
	RegDate time.Time
}

func (Student) TableName() string {
	return "tbl_student"
}


type Pet struct {
	ID uint64
	Name string
	Age uint8
}

func (Pet) TableName() string {
	return "tbl_pet"
}