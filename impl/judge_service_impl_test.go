package impl_test

import (
	"context"
	"io/ioutil"
	"judger/impl"
	util "judger/util"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestJudge(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(dir)
}

func TestCompile(t *testing.T) {
	path := util.GetPath()
	log.Println(path)
	sourceFile, err := os.Open(path + "/temp_data/test/aaa.c")
	if err != nil {
		log.Fatal(err.Error())
	}
	codeText, err := ioutil.ReadAll(sourceFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	//log.Printf("\n%s\n", codeText)

	err = impl.Compile(context.Background(), codeText, "c", path, "233323")
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("compile success")
	}

}
