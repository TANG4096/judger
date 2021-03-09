package util_test

import (
	"judger/util"
	"log"
	"testing"
)

func TestGetPath(t *testing.T) {
	log.Println("aaa")
	s := util.GetPath()
	log.Println(s)
}
