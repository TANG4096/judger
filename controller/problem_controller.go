package controller

import (
	"fmt"
	"judger/cache"
	"judger/model"
	"judger/pb"
	"judger/task"
	"judger/util"
	"strconv"

	"github.com/gin-gonic/gin"
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

func ProblemData(c *gin.Context) {
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
	task := task.GetSubmitTask()
	fmt.Println(req)
	fmt.Println(key)
	err = task.AddTask(key, req)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	util.Success(c, key)
}

func Judgestatus(c *gin.Context) {
	key, _ := c.GetQuery("id")
	result := cache.GetJudgeResuCache().Get(key)
	if result == "" {
		util.Success(c, "pending")
	} else {
		util.Success(c, result)
	}
}

func JudgestatusList(c *gin.Context) {
	uid := util.Query2Int(c, "uid")
	problemID := util.Query2Int(c, "problemID")
	contest := c.Query("contest")
	var contestID int = -1
	if contest != "" {
		contestID, _ = strconv.Atoi(contest)

	}
	list, err := model.GetJudgeStatusDataList(uid, problemID, contestID)
	if err != nil {
		util.CatchErr(c, err)
		return
	}
	util.Success(c, list)
}

func JudgestatusCode(c *gin.Context) {
	key := c.Query("id")
	result := cache.GetJudgeResuCache().Get(key)
	if result == "" {
		util.Success(c, "pending")
	} else {
		util.Success(c, result)
	}
}
