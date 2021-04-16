package impl_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"judger/impl"
	"judger/model"
	"judger/pb"
	util "judger/util"
	"log"
	"os"
	"testing"
	"time"
)

func TestJudge(t *testing.T) {
	server := impl.JudgeServiceServer{}
	req := &pb.JudgeRequest{
		ProblemID: 3,
		Type:      "c++",
		IsUpdate:  false,
	}
	path := util.GetPath() + "/temp_data/test/"
	sourceFile, err := ioutil.ReadFile(path + "a.cpp")
	if err != nil {
		fmt.Println(err)
	}
	req.SourceCode = sourceFile
	res, err := server.Judge(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", res)
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

func TestRun(t *testing.T) {
	path := util.GetPath() + "/temp_data/"
	fileName := "a+b"
	testdata := []model.ProblemTestData{
		{
			ProblemID: 1,
			Input:     []byte("1 2\n"),
			Ans:       []byte("3\n"),
		},
		{
			ProblemID: 1,
			Input:     []byte("1023 -1\n"),
			Ans:       []byte("1022\n"),
		},
	}
	res, err := impl.Run(context.Background(), path, fileName, testdata, 2, 256*(1<<20), 0)
	if err != nil {
		fmt.Println("err: " + err.Error())
	}
	fmt.Println(*res)
	time.Sleep(3 * time.Second)
}
