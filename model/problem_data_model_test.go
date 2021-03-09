package model_test

import (
	"fmt"
	"judger/model"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestInit(t *testing.T) {
	data := model.ProblemData{}
	fmt.Println(data)
}

func TestProblemDataInsert(t *testing.T) {
	data := model.ProblemData{
		Name:        "a+b",
		TimeLimit:   2000,
		MemoryLimit: 512 * (1 << (20)),
		JudgerID:    0,
	}
	data.Insert()
	res := model.ProblemData{Name: "a+b"}
	fmt.Printf("%v", res)
}
