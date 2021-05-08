package service

import (
	"judger/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context, data model.UserData) (gin.H, error) {
	h := gin.H{
		"errCode": 0,
		"msg":     "",
		"data":    nil,
	}
	err := data.Insert()

	if err != nil {
		errs := err.Error()
		if errs == "Error 1062: Duplicate entry 'test' for key 'user_name'" {
			h["errCode"] = 1062
			h["msg"] = "user name error"
		} else {
			h["errCode"] = 1000
			h["msg"] = errs
		}
	}
	return h, err
}
