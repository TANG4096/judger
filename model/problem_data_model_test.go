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
		Name:        "a+b",
		TimeLimit:   2,
		MemoryLimit: 512 * (1 << (20)),
		JudgerID:    0,
		Author:      "ty",
		Content:     ("输入a和b输出两个数字的和\n"),
		Input:       ("第一行一个整数n,表示数据组数，后面n行每行是要求和的两个数\n"),
		Output:      ("输出和，每行一个整数\n"),
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
