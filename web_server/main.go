// main.go
package main

import (
	"judger/controller"
	"judger/task"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sta-golang/go-lib-utils/log"
)

func register(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{ // 返回一个JSON，状态码是200，gin.H是map[string]interface{}的简写
			"message": "pong",
		})
	})
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/userdata/", controller.UserData)
	r.GET("/problem/", controller.Problem)
	r.GET("/problemdata/", controller.ProblemDataList)
	r.GET("/problemtag/", controller.ProblemTagList)
	r.GET("/sample/", controller.GetSample)
	r.POST("/submit/", controller.Submit)
	r.GET("/judgestatus/", controller.Judgestatus)
	r.GET("/judgestatuslist/", controller.JudgestatusList)
	r.GET("/judgestatuscode/", controller.JudgestatusCode)
	r.GET("/settingboard/", controller.SettingBoard)
	r.GET("/statusdatalist/", controller.StatusDataList)
	r.GET("/contestlist/", controller.GetContestList)
	r.POST("/addcontest/", controller.AddContest)
	r.POST("/addcontestproblem/", controller.AddContestProblem)
	r.GET("/currenttime/", controller.CurrentTime)
	r.GET("/contestinfo/", controller.ContestInfo)
	r.GET("/contestannouncement/", controller.ContestAnnouncement)
	r.GET("/contestregister/",controller.ContestRegister)
	return r
}

func LoggerToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		//日志格式
		log.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func main() {
	gin.Logger()
	task := task.GetSubmitTask()
	go task.Run()
	r := gin.Default() // 使用默认中间件（logger和recovery）
	r.Use(LoggerToFile())
	r = register(r)

	r.Run(":80") // 启动服务，并默认监听8080端口
}
