package model

import (
	"judger/db"

	"gorm.io/gorm"
)

type JudgeStatusData struct {
	gorm.Model
	Key       string `gorm:"index"`
	Uid       int    `gorm:"index"`
	UserName  string ``
	ProblemID int    `gorm:"index"`
	Result    string `gorm:"not null"`
	Language  string `gorm:"not null"`
	ContestID int    `gorm:"not null"`
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&JudgeStatusData{})
}

func (data *JudgeStatusData) Insert() error {
	return db.GetDB().Create(data).Error
}

func GetJudgeStatusDataList(temp *JudgeStatusData) (res []JudgeStatusData, err error) {
	res = []JudgeStatusData{}
	err = db.GetDB().Where(temp).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetStatusDataList(temp *JudgeStatusData, limit, offset int) (data []JudgeStatusData, err error) {
	data = []JudgeStatusData{}
	err = db.GetDB().Where(temp).Offset(offset).Limit(limit).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetStatusDataCnt(temp *JudgeStatusData) (cnt *int64, err error) {
	cnt = new(int64)
	err = db.GetDB().Model(&JudgeStatusData{}).Where(temp).Count(cnt).Error
	if err != nil {
		return nil, err
	}
	return cnt, nil
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
