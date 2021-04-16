package util_test

import (
	"fmt"
	"judger/util"
	"testing"
	"time"

	"github.com/sta-golang/go-lib-utils/log"
)

func TestLogInit(t *testing.T) {
	fmt.Println(util.GetPath())
	log.Info("test")
	log.Debug("debug")
	log.Info("233342")
	time.Sleep(5 * time.Second)
}
