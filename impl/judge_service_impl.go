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
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/sta-golang/go-lib-utils/log"
)

const outBufferSize int = 20 * (1 << 20)

type JudgeServiceServer struct {
}

func (ju *JudgeServiceServer) Judge(ctx context.Context, req *pb.JudgeRequest) (res *pb.JudgeResponse, err error) {
	dir := util.GetPath()
	fileName := fmt.Sprintf("%d_%d_%d", req.ProblemID, req.UserID, time.Now().Unix())
	err = Compile(ctx, req.SourceCode, req.Type, dir, fileName)
	if err != nil {
		if err.Error() == "compile error" {
			log.Infof("file %s comple error", fileName)
			return &pb.JudgeResponse{Response: "compile error"}, nil
		} else {
			util.LogErr(err)
			return nil, err
		}
	}
	log.Infof("%s compile complete", fileName)
	testDataList, err := GetTestDataList(uint(req.ProblemID), req.IsUpdate)
	if err != nil {
		util.LogErr(err)
		return nil, err
	}
	problemData := model.ProblemData{}
	problemData.ID = uint(req.ProblemID)
	err = problemData.GetLimit()
	if err != nil {
		return nil, err
	}
	result, err := Run(ctx, dir+"/temp_data/", fileName, testDataList, problemData.TimeLimit, problemData.MemoryLimit, problemData.JudgerID)
	if err != nil {
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
		err := ioutil.WriteFile(fileName+suffix, codeText, os.FileMode(777))
		if err != nil {
			return err
		}
		compilerNameMap := util.GetConfigMap("complierName")
		//log.Println(compilerNameMap)
		cmd := exec.Command(compilerNameMap[Type], "-o", fileName, fileName+suffix)
		out, err := cmd.CombinedOutput()
		if err != nil {
			if len(out) != 0 {
				return errors.New("compile error")
			}
			return err
		}

	default:
		return errors.New("not support this language")
	}
	return nil
}

func Run(ctx context.Context, dir, fileName string, testDataList []model.ProblemTestData, TimeLimit, MemoryLimit, judgerID uint) (res *string, err error) {
	var ans string
	for caseID, testData := range testDataList {
		outFile, err := ioutil.TempFile(dir, "out")
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(dir + fileName)
		cmd.Stdout = outFile
		inWriter, err := cmd.StdinPipe()
		if err != nil {
			return nil, err
		}
		r1, _, errno := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if errno != 0 {
			return nil, err
		}
		res := int(r1)
		if res == 0 {
			defer outFile.Close()
			err := syscall.Setrlimit(syscall.RLIMIT_CPU, &syscall.Rlimit{Cur: uint64(TimeLimit), Max: uint64(TimeLimit)})
			if err != nil {
				return nil, err
			}
			err = syscall.Setrlimit(syscall.RLIMIT_AS, &syscall.Rlimit{Cur: uint64(MemoryLimit), Max: uint64(MemoryLimit)})
			if err != nil {
				return nil, err
			}
			err = syscall.Setrlimit(syscall.RLIMIT_DATA, &syscall.Rlimit{Cur: uint64(MemoryLimit), Max: uint64(MemoryLimit)})
			if err != nil {
				return nil, err
			}
			cmd.Start()
			n, err := inWriter.Write(testData.Input)
			fmt.Printf("write %d bytes to stdInput\n", n)
			if err != nil {
				return nil, err
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
					return nil, err
				}
			}
			os.Exit(0)
		} else {
			process, err := os.FindProcess(res)
			if err != nil {
				return nil, err
			}
			ps, err := process.Wait()
			if err != nil {
				return nil, err
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
					stat, err := outFile.Stat()
					if err != nil {
						return nil, err
					}
					fmt.Println("outfile size: ", stat.Size())
					out, err := ioutil.ReadFile(outFile.Name())
					if err != nil {
						return nil, err
					}
					log.Infof("Problem %d case %d output %d bytes", testData.ProblemID, caseID+1, len(out))
					fmt.Println(out)
					fmt.Println(testData.Ans)
					if !bytes.Equal(out, testData.Ans) {
						ans = "wrong answer"
						break
					}
				} else {
					//TODO special judger
				}
			}
			os.Remove(outFile.Name())
		}
	}
	if ans == "" {
		ans = "accept"
	}
	return &ans, nil
}
