package model

import (
	"judger/db"

	"gorm.io/gorm"
)

type JudgeStatusData struct {
	gorm.Model
	Key       string `gorm:"index"`
	Uid       int    `gorm:"index"`
	ProblemID int    `gorm:"index"`
	Result    int    `gorm:"not null"`
	Language  string `gorm:"not null"`
	ContestID int    `gorm:"not null"`
}

func (data *JudgeStatusData) Insert() error {
	return db.GetDB().Create(data).Error
}

func GetJudgeStatusDataList(uid, problemID, contestID int) (res []JudgeStatusData, err error) {
	temp := &JudgeStatusData{}

	temp.Uid = uid
	temp.ProblemID = problemID
	if contestID != -1 {
		temp.ContestID = contestID
	}
	err = db.GetDB().Where(temp).Find(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetJudgeStatus(key string) (res *JudgeStatusData, err error) {
	res = &JudgeStatusData{
		Key: key,
	}
	err = db.GetDB().Where(res).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
