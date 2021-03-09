package model_test

import (
	"fmt"
	"judger/model"
	"testing"
)

func TestPoblemTestDataInsert(t *testing.T) {
	data := model.ProblemTestData{
		ProblemID: 3,
		Input:     []byte("2\n1 2\n2 3\n"),
		Ans:       []byte("3\n5\n"),
	}
	err := data.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}
	dataList, err := model.GetTestDataList(3)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v\n%s\n%s\n", dataList, dataList[0].Input, dataList[0].Ans)
}
