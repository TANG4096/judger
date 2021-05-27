package model

import (
	"judger/db"

	"gorm.io/gorm"
)

type SubmitCode struct {
	gorm.Model
	Key      string `gorm:"index"`
	Code     string
	Language string
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&SubmitCode{})
}
func (data *SubmitCode) Insert() error {
	return db.GetDB().Create(data).Error
}

func GetSubmitCode(key string) (data *SubmitCode, err error) {
	data = &SubmitCode{}
	err = db.GetDB().Where(&SubmitCode{Key: key}).First(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
