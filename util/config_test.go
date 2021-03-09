package util_test

import (
	"judger/util"
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	log.Println(util.DbConfigMap)
	mp := util.GetDbConfigMap("test")
	log.Println(mp)
}
