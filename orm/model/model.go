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

type StudentScore struct {
	ID uint64
	Course string
	StuNo string
	Score uint
}

func (StudentScore) TableName() string {
	return "tbl_student_score"
}

type Pet struct {
	ID uint64
	Name string
	Age uint8
}

func (Pet) TableName() string {
	return "tbl_pet"
}