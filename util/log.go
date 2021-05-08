package util

import (
	"time"

	"github.com/sta-golang/go-lib-utils/log"
)

func init() {

	//---------------------------
	alone := []string{log.GetLevelName(log.INFO), log.GetLevelName(log.ERROR)}
	logConfig := &log.FileLogConfig{
		FileDir:     GetPath() + "/log/judge_server",
		FileName:    "judge_server",
		DayAge:      7,
		LogLevel:    log.INFO,
		MaxSize:     0,
		AloneWriter: alone,
		Prefix:      "judger",
	}
	Logger := log.NewFileLogAndAsync(logConfig, time.Second*3)
	log.SetGlobalLogger(Logger)
}

func LogErr(err error) {
	log.Errorf("err: %v", err)
}
