package controller

import (
	"errors"
	"fmt"
	"judger/model"
	"judger/service"
	"judger/util"

	"github.com/gin-gonic/gin"
	"github.com/sta-golang/go-lib-utils/log"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	postForm, err := util.GetPostForm(c)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	data := model.UserData{
		UserName: postForm.GetValue2String("username"),
		PassWord: postForm.GetValue2String("password"),
		Name:     postForm.GetValue2String("name"),
		Email:    postForm.GetValue2String("email"),
	}
	fmt.Println(data)
	res, err := service.Register(c, data)
	if err != nil {
		log.Error(err)
	}
	c.JSON(200, res)
}

func Login(c *gin.Context) {

	postForm, err := util.GetPostForm(c)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	UserName := postForm.GetValue2String("username")
	PassWord := postForm.GetValue2String("password")
	temp, err := model.GetUserDataByUserName(UserName)
	fmt.Printf("username: %s\npassword: %s\n", UserName, PassWord)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			util.Failed(c, "用户名不存在")
		} else {
			util.CatchErr(c, err)
		}
		return
	}
	if temp.PassWord != PassWord {
		util.Failed(c, "密码错误")
	} else {
		res := gin.H{
			"username": temp.UserName,
			"name":     temp.Name,
			"uid":      temp.ID,
			"type":     temp.Type,
		}
		util.Success(c, res)
	}
}

func UserData(c *gin.Context) {

}
func SettingBoard(c *gin.Context) {
	c.JSON(200, []string{})

}
