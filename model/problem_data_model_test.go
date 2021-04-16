package model_test

import (
	"fmt"
	"judger/db"
	"judger/model"
	"log"
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
		JudgerID:    1,
	}
	data.Insert()
	res := model.ProblemData{}
	err := db.GetDB().Find(&model.ProblemData{Name: "a+b"}).First(&res).Error
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("%v", res)
}
