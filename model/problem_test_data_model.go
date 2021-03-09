package model

import (
	"judger/db"

	"github.com/jinzhu/gorm"
)

type ProblemTestData struct {
	gorm.Model
	ProblemID uint   `gorm: "index"`
	Input     []byte `gorm: "NOT_NULL";`
	Ans       []byte `gorm: "NOT_NULL"`
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&ProblemTestData{})
}

func GetTestDataList(problemID uint) (dataList []ProblemTestData, err error) {
	db := db.GetDB()
	if err = db.Where(&ProblemTestData{ProblemID: problemID}).Find(&dataList).Error; err != nil {
		return nil, err
	}
	return dataList, err
}

func (data *ProblemTestData) Insert() error {
	return db.GetDB().Create(data).Error
}
