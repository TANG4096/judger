package model

import (
	"judger/db"
	"time"

	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	Name      string `gorm:"unique;index"`
	BeginTime time.Time
	LastTime  time.Time
	Type      string
	Auth      int
	Creator   string
	Des       string
	Note      string
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&Contest{})
}

func (data *Contest) Insert() (id *uint, err error) {
	id = new(uint)
	err = db.GetDB().Create(data).Where(data).First(data).Error
	if err != nil {
		return nil, err
	}
	id = &data.ID
	return
}

func GetContestList(temp *Contest, offset, limit int) (data []Contest, err error) {
	data = []Contest{}
	err = db.GetDB().Where(temp).Offset(offset).Limit(limit).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetContest(id uint) (data *Contest, err error) {
	err = db.GetDB().Where("id=?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return
}


