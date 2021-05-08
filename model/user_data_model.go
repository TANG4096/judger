package model

import (
	"errors"
	"judger/db"

	"gorm.io/gorm"
)

type UserData struct {
	gorm.Model
	UserName string `gorm:"unique;index"`
	PassWord string `gorm:"NOT NULL"`
	Name     string `gorm:"NOT NULL"`
	Email    string `gorm:"NOT NULL"`
	Type     string `gorm:"default:1"`
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&UserData{})
}

func (data *UserData) Insert() error {
	return db.GetDB().Create(data).Error
}

func GetUserDataByUserName(userName string) (data *UserData, err error) {
	data = &UserData{}
	if userName == "" {
		return nil, errors.New("empty parameter")
	}
	err = db.GetDB().Where(&UserData{UserName: userName}).First(&data).Error
	return data, err
}
