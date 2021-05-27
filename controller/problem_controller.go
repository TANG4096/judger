package controller

import (
	"errors"
	"fmt"
	"judger/cache"
	"judger/model"
	"judger/pb"
	"judger/task"
	"judger/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sta-golang/go-lib-utils/log"
	"gorm.io/gorm"
)

func Problem(c *gin.Context) {
	id := util.Query2Int(c, "id")
	data, err := model.GetProblemData(id)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	util.Success(c, data)
}

func ProblemDataList(c *gin.Context) {
	limit := util.Query2Int(c, "limit")
	offset := util.Query2Int(c, "offset")
	auth := util.Query2Int(c, "auth")

	problemDataList, err := model.GetProblemList(limit, offset, auth)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	cnt, err := model.GetProbemCount()
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	res := gin.H{
		"results": problemDataList,
		"count":   cnt,
	}
	//fmt.Println(problemDataList)
	util.Success(c, res)
}

func ProblemTagList(c *gin.Context) {

	util.Success(c, gin.H{})
}

func GetSample(c *gin.Context) {
	id := util.Query2Uint(c, "id")
	dataList, err := model.GetSample(id)
	if err != nil {
		util.CatchErr(c, err)
		fmt.Println(err)
		return
	}
	InputList := []string{}
	OutputList := []string{}
	for _, v := range dataList {
		InputList = append(InputList, string(v.Input))
		OutputList = append(OutputList, string(v.Ans))
	}
	res := gin.H{
		"input":  InputList,
		"output": OutputList,
	}
	//fmt.Println(res)
	util.Success(c, res)
}

func Submit(c *gin.Context) {
	postFrom, err := util.GetPostForm(c)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	fmt.Println(postFrom)

	req := &pb.JudgeRequest{
		ProblemID:  int32(postFrom.GetValue2Int("problemID")),
		UserID:     int32(postFrom.GetValue2Int("uid")),
		Type:       postFrom.GetValue2String("language"),
		SourceCode: []byte(postFrom.GetValue2String("code")),
	}
	key := strconv.Itoa(int(req.ProblemID)) + "_" + strconv.Itoa(int(req.UserID)) + "_" + postFrom.GetValue2String("submittime")
	code := model.SubmitCode{
		Key:      key,
		Code:     string(req.SourceCode),
		Language: req.Type,
	}
	err = code.Insert()
	if err != nil {
		util.CatchErr(c, err)
		return
	}

	task := task.GetSubmitTask()
	fmt.Println(req)
	fmt.Println(key)
	userName := postFrom.GetValue2String("user")
	err = task.AddTask(userName, key, req)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	util.Success(c, key)
}

func Judgestatus(c *gin.Context) {
	key, _ := c.GetQuery("id")
	result := cache.GetJudgeResuCache().Get(key)
	if result == nil {
		result, err := model.GetJudgeStatus(key)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				util.Success(c, "pending")
				return
			}
			util.CatchErr(c, err)
			return
		}
		util.Success(c, result.Result)
	} else {
		util.Success(c, result.Result)
	}
}

func JudgestatusList(c *gin.Context) {

	user := c.Query("user")
	problemID, ok := util.GetQuery2Int(c, "problemID")
	contest, ok := c.GetQuery("contest")
	var contestID int = 0
	if ok {
		contestID, _ = strconv.Atoi(contest)
	}
	temp := &model.JudgeStatusData{
		UserName:  user,
		ContestID: contestID,
		ProblemID: problemID,
	}
	list, err := model.GetJudgeStatusDataList(temp)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	util.Success(c, list)
}

func StatusDataList(c *gin.Context) {

	limit := util.Query2Int(c, "limit")
	offset := util.Query2Int(c, "offset")
	temp := model.JudgeStatusData{}
	uid, ok := util.GetQuery2Int(c, "uid")
	if ok {
		temp.Uid = uid
	}
	language, ok := c.GetQuery("language")
	if ok {
		temp.Language = language
	}
	result, ok := c.GetQuery("result")
	if ok {
		temp.Result = result
	}
	contest, ok := util.GetQuery2Int(c, "contest")
	if ok {
		temp.ContestID = contest
	}
	data, err := model.GetStatusDataList(&temp, limit, offset)
	if err != nil {
		log.Error(err)
		util.Failed(c, err.Error())
		return
	}
	cnt, err := model.GetStatusDataCnt(&temp)
	if err != nil {
		log.Error(err)
		util.Failed(c, err.Error())
		return
	}
	res := gin.H{
		"results": data,
		"count":   *cnt,
	}
	util.Success(c, res)
}

func JudgestatusCode(c *gin.Context) {
	key := c.Query("id")
	data, err := model.GetSubmitCode(key)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	util.Success(c, data)
}
