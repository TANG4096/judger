package controller

import (
	"judger/model"
	"judger/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sta-golang/go-lib-utils/log"
)

func GetContestList(c *gin.Context) {
	limit := util.Query2Int(c, "limit")
	offset := util.Query2Int(c, "offset")
	req := &model.Contest{}

	list, err := model.GetContestList(req, offset, limit)
	if err != nil {
		log.Error(err)
		util.Failed(c, err)
		return
	}
	util.Success(c, gin.H{
		"results": list,
		"count":   0,
	})
}

func AddContest(c *gin.Context) {
	postFrom, err := util.GetPostForm(c)
	if err != nil {
		log.Error(err)
		util.Failed(c, err)
		return
	}
	beginTime, err := time.Parse("2006-01-02 15:04:05", postFrom.GetValue2String("begintime"))
	if err != nil {
		log.Error(err)
		util.Failed(c, err)
		return
	}
	lastTime, err := time.Parse("2006-01-02 15:04:05", postFrom.GetValue2String("lasttime"))
	if err != nil {
		log.Error(err)
		util.Failed(c, err)
		return
	}
	data := &model.Contest{
		Creator:   postFrom.GetValue2String("creator"),
		Name:      postFrom.GetValue2String("title"),
		Des:       postFrom.GetValue2String("des"),
		Note:      postFrom.GetValue2String("note"),
		Type:      postFrom.GetValue2String("type"),
		Auth:      postFrom.GetValue2Int("auth"),
		BeginTime: beginTime,
		LastTime:  lastTime,
	}
	id, err := data.Insert()
	if err != nil {
		log.Error(err)
		util.Failed(c, err)
		return
	}
	util.Success(c, gin.H{"id": *id})
}

func AddContestProblem(c *gin.Context) {
	postFrom, err := util.GetPostForm(c)
	if err != nil {
		log.Error(err)
		util.Failed(c, err)
		return
	}
	data := model.ContestProblem{
		ContestID:   uint(postFrom.GetValue2Int("contestid")),
		ProblemID:   uint(postFrom.GetValue2Int("problemid")),
		ProblemName: postFrom.GetValue2String("problemtitle"),
		Rank:        uint(postFrom.GetValue2Int("rank")),
	}
	err = data.Insert()
	if err != nil {
		log.Error(err)
		util.Failed(c, err)
		return
	}
	util.Success(c, nil)
}

func CurrentTime(c *gin.Context) {
	c.JSON(200, time.Now())
}

func ContestInfo(c *gin.Context) {
	id := util.Query2Uint(c, "id")
	data, err := model.GetContest(id)
	if err != nil {
		log.Error(err)
		util.Failed(c, err)
		return
	}
	util.Success(c, *data)
}

func ContestAnnouncement(c *gin.Context) {
	c.JSON(200, nil)
}

func ContestRegister(c *gin.Context){
	
}
