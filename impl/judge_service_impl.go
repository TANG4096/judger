package impl

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"judger/cache"
	"judger/model"
	"judger/pb"
	util "judger/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

const outBufferSize int = 20 * (1 << 20)

type JudgeServiceServer struct {
}

func (ju *JudgeServiceServer) Judge(ctx context.Context, req *pb.JudgeRequest) (res *pb.JudgeResponse, err error) {
	log.Printf("%v", req)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("获取运行路径失败 %v\n", err)
		return nil, err
	}
	fileName := fmt.Sprintf("%d_%d_%d", req.ProblemID, req.UserID, time.Now().Unix())

	err = Compile(ctx, req.SourceCode, req.Type, dir, fileName)
	if err != nil {
		if err.Error() == "compile error" {
			return &pb.JudgeResponse{Response: "compile error"}, nil
		} else {
			return nil, err
		}
	}
	testDataList, err := GetTestDataList(uint(req.ProblemID), req.IsUpdate)
	if err != nil {
		util.LogPrintln(err.Error())
		return nil, err
	}
	problemData := model.ProblemData{}
	problemData.ID = uint(req.ProblemID)
	err = problemData.GetLimit()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	result, err := Run(ctx, dir+"/temp_data/", fileName, testDataList, problemData.TimeLimit, problemData.MemoryLimit, problemData.JudgerID)
	if err != nil {
		util.LogPrintln(err.Error())
		return nil, err
	}
	return &pb.JudgeResponse{Response: *result}, err
}

func (ju *JudgeServiceServer) JudgeClientAndServerStream(ctx context.Context, req *pb.JudgeRequest) (res *pb.JudgeResponse, err error) {

	return res, err
}

func GetTestDataList(problemID uint, isUpdate bool) ([]model.ProblemTestData, error) {
	cache := cache.GetTestDataCache()
	testData := cache.Get(problemID)
	if isUpdate || len(testData) == 0 {
		testData, err := model.GetTestDataList(problemID)
		if err != nil {
			util.LogPrintln(err.Error())
			return nil, err
		}
		cache.Update(problemID, testData)
	}
	return testData, nil
}

func Compile(ctx context.Context, codeText []byte, Type, path, fileName string) error {
	switch Type {
	case "python":

	case "c", "c++":
		suffixMap := util.GetConfigMap("sourceFileSuffix")
		suffix := suffixMap[Type]
		fileName = path + "/temp_data/" + fileName
		log.Printf("%s\n", fileName)
		err := ioutil.WriteFile(fileName+suffix, codeText, os.FileMode(777))
		if err != nil {
			return err
		}
		compilerNameMap := util.GetConfigMap("complierName")
		//log.Println(compilerNameMap)
		cmd := exec.Command(compilerNameMap[Type], "-o", fileName, fileName+suffix)
		out, err := cmd.CombinedOutput()
		log.Printf("%s\n", out)
		if err != nil {
			if len(out) != 0 {
				return errors.New("compile error")
			}
			log.Println("complie fail")
			return err
		}
	default:
		return errors.New("not support this language")
	}
	return nil
}

func Run(ctx context.Context, dir, fileName string, testDataArr []model.ProblemTestData, TimeLimit, MemoryLimit, judgerID uint) (res *string, err error) {
	var ans string
	for _, testData := range testDataArr {
		outFile, err := ioutil.TempFile(dir, "out")
		if err != nil {
			log.Println(err.Error())
		}
		cmd := exec.Command(fileName)
		cmd.Stdout = outFile
		inWriter, err := cmd.StdinPipe()
		if err != nil {
			log.Println(err.Error())
		}
		r1, _, errno := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if errno != 0 {
			log.Println(err.Error())
		}
		res := int(r1)
		if res == 0 {

			err := syscall.Setrlimit(syscall.RLIMIT_CPU, &syscall.Rlimit{Cur: uint64(TimeLimit), Max: uint64(TimeLimit)})
			if err != nil {
				log.Println(err.Error())
			}
			err = syscall.Setrlimit(syscall.RLIMIT_AS, &syscall.Rlimit{Cur: uint64(MemoryLimit), Max: uint64(MemoryLimit)})
			if err != nil {
				log.Println(err.Error())
			}
			err = syscall.Setrlimit(syscall.RLIMIT_DATA, &syscall.Rlimit{Cur: uint64(MemoryLimit), Max: uint64(MemoryLimit)})
			if err != nil {
				log.Println(err.Error())
			}
			cmd.Start()
			n, err := inWriter.Write(testData.Input)
			fmt.Printf("write %d bytes\n", n)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = cmd.Wait()
			if err != nil {
				value, ok := (err).(*exec.ExitError)
				if ok {
					status := value.Sys().(syscall.WaitStatus)
					switch {
					case status.Signaled():
						os.Exit(int(status.Signal()))
					default:
						os.Exit(int(syscall.SIGUSR1))
					}
				} else {
					log.Println("err: ", err.Error())
				}
			}
			os.Exit(0)
		} else {
			fmt.Println("!!!!")
			process, err := os.FindProcess(res)
			if err != nil {
				log.Println(err.Error())
			}
			log.Println(process.Pid)
			ps, err := process.Wait()
			if err != nil {
				log.Println(err.Error())
			}

			code := syscall.Signal(ps.ExitCode())
			if code != 0 {
				fmt.Println("code: ", ps.ExitCode())
				switch code {
				case syscall.SIGSEGV:
					ans = "segment error"
				case syscall.SIGKILL:
					ans = "time limit error"
				case syscall.SIGUSR1:
					ans = "run time error"
				}
				break
			} else {
				if judgerID == 0 {
					out, err := ioutil.ReadAll(outFile)
					if err != nil {
						return nil, err
					}
					if !bytes.Equal(out, testData.Ans) {
						ans = "wrong answer"
						break
					}
				} else {
					//TODO special judger
				}
			}
			outFile.Close()
		}
	}
	if ans == "" {
		ans = "accept"
	}
	return &ans, nil
}
