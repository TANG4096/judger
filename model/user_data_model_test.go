package model_test

import (
	"fmt"
	"judger/db"
	"judger/model"
	"testing"

	_ "gorm.io/driver/mysql"
)

func TestUserDataInster(t *testing.T) {
	data := model.UserData{
		UserName: "test",
		PassWord: "123",
		Name:     "111",
		Email:    "18329676365@163.com",
	}
	err := data.Insert()
	if err != nil {
		fmt.Println("err: ", err)
	}
}

func TestGetUserDataByUserName(t *testing.T) {
	data, err := model.GetUserDataByUserName("")
	if err != nil {
		fmt.Println("err： ", err)
		return
	}
	fmt.Printf("%v", data)
}

func TestGetUpdatae(t *testing.T) {
	temp := &model.UserData{
		UserName: "test",
	}
	data := &model.UserData{
		Type: "3",
	}
	err := db.GetDB().Model(&model.UserData{}).Where(temp).Updates(data).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%v", data)
}
