package model

import (
	"judger/db"

	"github.com/jinzhu/gorm"
)

type ProblemData struct {
	gorm.Model
	Name        string `gorm: "index"`
	TimeLimit   uint
	MemoryLimit uint
	JudgerID    uint
	AcceptCnt   int
	SubmitCnt   int
	Content     []byte
	Extend      []byte
}

func init() {
	db := db.GetDB()
	db.AutoMigrate(&ProblemData{})
}

func (data *ProblemData) GetLimit() (err error) {
	db := db.GetDB()
	if err = db.Where(data, "time_limit", "memory_limit", "judger_id").First(&data).Error; err != nil {
		return err
	}
	return nil
}

func (data *ProblemData) Insert() error {
	return db.GetDB().Create(data).Error
}
