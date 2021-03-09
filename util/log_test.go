package util_test

import (
	"judger/util"
	"testing"
)

func TestLogFatalln(t *testing.T) {
	util.LogFatalln("aa")
}
