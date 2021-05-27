package model_test

import (
	"fmt"
	"judger/db"
	"judger/model"
	"testing"
)

func TestXXX(t *testing.T) {
	db := db.GetDB()
	db.AutoMigrate(&model.Contest{})
}

func TestGetContestList(t *testing.T) {
	list, err := model.GetContestList(&model.Contest{}, 0, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
