package model

import (
	"judger/db"

	"gorm.io/gorm"
)

type ProblemData struct {
	gorm.Model
	Name        string `gorm:"index"`
	TimeLimit   uint
	MemoryLimit uint
	JudgerID    uint
	AcceptCnt   int `gorm:"Default:1"`
	SubmitCnt   int `gorm:"Default:1"`
	MleCnt      int `gorm:"Default:1"`
	TleCnt      int `gorm:"Default:1"`
	RteCnt      int `gorm:"Default:1"`
	PeCnt       int `gorm:"Default:1"`
	Auth        int `gorm:"Default:1"`
	Level       int `gorm:"Default:1"`
	Author      string
	Content     string
	Input       string
	Output      string
	Hint        string
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&ProblemData{})
}

func (data *ProblemData) GetLimit() (err error) {
	db := db.GetDB()
	if err = db.Select("time_limit", "memory_limit", "judger_id").Where(data).First(&data).Error; err != nil {
		return err
	}
	return nil
}

func (data *ProblemData) Insert() error {
	return db.GetDB().Create(data).Error
}

func GetProblemList(limit, offset, auth int) ([]ProblemData, error) {
	ans := []ProblemData{}
	err := db.GetDB().Limit(limit).Offset(offset).Where("auth <= ?", auth).Find(&ans).Error
	if err != nil {
		return nil, err
	}
	return ans, err
}

func GetProblemData(id int) (*ProblemData, error) {
	ans := ProblemData{}
	err := db.GetDB().First(&ans).Error
	if err != nil {
		return nil, err
	}
	return &ans, err
}

func GetProbemCount() (int64, error) {
	var cnt int64
	err := db.GetDB().Model(&ProblemData{}).Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return cnt, err
}
