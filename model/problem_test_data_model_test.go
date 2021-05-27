package model_test

import (
	"fmt"
	"judger/model"
	"testing"
)

func TestPoblemTestDataInsert(t *testing.T) {
	data := model.ProblemTestData{
		ProblemID: 1,
		Input:     []byte("2\n2 2\n2 3\n"),
		Ans:       []byte("4\n5\n"),
	}
	fmt.Printf("插入数据为：\n%v\n", data)
	err := data.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("数据插入成功")
	dataList, err := model.GetTestDataList(3)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v\n", dataList)
}

func TestGetSample(t *testing.T) {
	data, err := model.GetSample(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func TestGetTestDataList(t *testing.T) {
	list, err := model.GetTestDataList(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
