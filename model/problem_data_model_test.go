package model_test

import (
	"fmt"
	"judger/model"
	"log"
	"testing"

	_ "gorm.io/driver/mysql"
)

func TestInit(t *testing.T) {
	data := model.ProblemData{}
	fmt.Println(data)
}

func TestProblemDataInsert(t *testing.T) {
	data := model.ProblemData{
		Name:        "异或加密",
		TimeLimit:   2,
		MemoryLimit: 512 * (1 << (20)),
		JudgerID:    0,
		Author:      "ty",
	}
	data.Insert()
	list, err := model.GetProblemList(10, 0, 1)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("%v", list)
}

func TestGetproblemCount(t *testing.T) {
	cnt, err := model.GetProbemCount()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cnt)
}
