package model

import (
	"judger/db"

	"gorm.io/gorm"
)

type ContestProblem struct {
	gorm.Model
	ContestID   uint `gorm:"index"`
	ProblemID   uint
	ProblemName string
	Rank        uint
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&ContestProblem{})
}

func (data *ContestProblem) Insert() error {
	return db.GetDB().Create(data).Error
}
