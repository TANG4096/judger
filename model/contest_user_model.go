package model

import (
	"judger/db"

	"gorm.io/gorm"
)

type ContestUser struct {
	gorm.Model
	ContestID uint
	UserName  string
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&ContestUser{})
}

func (data *ContestUser) Insert() error {
	return db.GetDB().Create(data).Error
}

func GetContestUserList() {

}
