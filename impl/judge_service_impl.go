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
	t1 := time.Now()
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
	CompileTime := time.Since(t1)
	log.Infof("%s compile complete", fileName)
	fmt.Printf("编译用时：%v\n", CompileTime)
	t1 = time.Now()
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
	d := time.Since(t1)
	fmt.Printf("拉取题目信息用时：%v\n", d)
	result, err := Run(ctx, dir+"/temp_data/", fileName, testDataList, problemData.TimeLimit, problemData.MemoryLimit, problemData.JudgerID)
	if err != nil {
		return nil, err
	}

	return &pb.JudgeResponse{Response: *result}, err
}

func (ju *JudgeServiceServer) JudgeClientAndServerStream(ctx context.Context, req *pb.JudgeRequest) (res *pb.JudgeResponse, err error) {

	return res, err
}

func GetTestDataList(problemID uint, isUpdate bool) (testData []model.ProblemTestData, err error) {
	cache := cache.GetTestDataCache()
	testData = cache.Get(problemID)
	path := util.GetPath()
	if len(testData) == 0 {
		t := 1
		for {
			//fmt.Println(t)
			filePath := fmt.Sprintf("%s/temp_data/test_case/%d/", path, int(problemID))
			inputFileName := fmt.Sprintf("%s%d.in", filePath, t)
			outputFileName := fmt.Sprintf("%s%d.ans", filePath, t)
			data := model.ProblemTestData{ProblemID: problemID}
			data.Input, err = ioutil.ReadFile(inputFileName)
			if err != nil {
				//fmt.Println(err)
				break
			}
			data.Ans, err = ioutil.ReadFile(outputFileName)
			if err != nil {
				//fmt.Println(err)
				break
			}
			//data.Input = crlf2lf(data.Input)
			data.Ans = crlf2lf(data.Ans)
			testData = append(testData, data)
			t++
		}
	}
	if len(testData) != 0 {
		return testData, nil
	}
	if isUpdate {
		testData, err = model.GetTestDataList(problemID)
		if err != nil {
			return nil, err
		}
		cache.Update(problemID, testData)
	}
	return testData, nil
}

func Compile(ctx context.Context, codeText []byte, Type, path, fileName string) error {
	switch Type {
	case "Python":

	case "C", "C++":
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
		defer os.Remove(fileName + suffix)
		out, err := cmd.CombinedOutput()
		if err != nil {
			if len(out) != 0 {
				return errors.New("Compile Error")
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
	//fmt.Println(testDataList)

	defer os.Remove(dir + fileName)
	for caseID, testData := range testDataList {
		fmt.Println("test case: ", caseID+1)
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
				os.Exit(int(syscall.SIGUSR1))
			}
			err = syscall.Setrlimit(syscall.RLIMIT_AS, &syscall.Rlimit{Cur: uint64(MemoryLimit), Max: uint64(MemoryLimit)})
			if err != nil {
				os.Exit(int(syscall.SIGUSR1))
			}
			err = syscall.Setrlimit(syscall.RLIMIT_DATA, &syscall.Rlimit{Cur: uint64(MemoryLimit), Max: uint64(MemoryLimit)})
			if err != nil {
				os.Exit(int(syscall.SIGUSR1))
			}
			cmd.Start()
			n, err := inWriter.Write(testData.Input)
			fmt.Printf("write %d bytes to stdInput\n", n)
			if err != nil {
				fmt.Println(err)
				os.Exit(int(syscall.SIGUSR1))
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
					os.Exit(int(syscall.SIGUSR1))
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
			runTime := ps.SystemTime() + ps.UserTime()
			fmt.Printf("运行用时：%v\n", runTime)
			code := syscall.Signal(ps.ExitCode())
			if code != 0 {
				log.Infof("code: %d\n", ps.ExitCode())
				fmt.Println("code: ", ps.ExitCode())
				switch code {
				case syscall.SIGSEGV:
					ans = "Memory Limit Exceeded"
				case syscall.SIGKILL:
					ans = "Time Limit Exceeded"
				case syscall.SIGUSR1:
					ans = "Runtime Rrror"
				}
				return &ans, nil
			} else {
				t1 := time.Now()
				fmt.Println("judgeID: ", judgerID)
				stat, err := outFile.Stat()
				if err != nil {
					return nil, err
				}
				log.Infof("outfile size: %d", stat.Size())
				out, err := ioutil.ReadFile(outFile.Name())
				if err != nil {
					return nil, err
				}
				log.Infof("Problem %d case %d output %d bytes", testData.ProblemID, caseID+1, len(out))
				//fmt.Printf("out:\n%v", out)
				//fmt.Printf("ans:\n%v", testData.Ans)

				if judgerID == 0 {
					if !bytes.Equal(out, testData.Ans) {
						ans = "Wrong Answer"
						return &ans, nil
					}
				} else {
					if !SpecialJudge(judgerID, out, testData.Ans) {
						ans = "Wrong Answer"
						return &ans, nil
					}
				}
				CompareTime := time.Since(t1)
				fmt.Printf("结果比较用时：%v\n", CompareTime)
			}
			os.Remove(outFile.Name())
		}
	}
	ans = "Accept"
	return &ans, nil
}

func crlf2lf(s []byte) []byte {
	cnt := 0
	book := make([]bool, len(s))
	for k, v := range s {
		if v == 13 {
			book[k] = true
			cnt++
		} else {
			if cnt != 0 {
				s[k-cnt] = v
			}

		}
	}
	return s[:len(s)-cnt]
}

func SpecialJudge(SpecialJudgeID uint, out, ans []byte) bool {
	return true
}
