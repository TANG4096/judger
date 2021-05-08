package task

import (
	"fmt"
	"judger/cache"
	"judger/pb"
	"judger/service"
	"runtime"
	"time"

	"github.com/sta-golang/go-lib-utils/log"
)

type SubmitTask struct {
	bufferQueue *LKQueue
	pool        *TaskPool
}
type SubmitParam struct {
	key string
	req *pb.JudgeRequest
}

func GetSubmitTask() *SubmitTask {
	task := SubmitTask{}
	task.bufferQueue = NewLKQueue()
	task.pool, _ = NewPool(100)
	return &task
}

func (t *SubmitTask) Run() {
	for {
		time.Sleep(1 * time.Second)
		res := t.bufferQueue.Dequeue()
		if res != nil {
			fmt.Println("res:", res)
			t.pool.Submit(func() error {
				param := res.(*SubmitParam)
				log.Infof("submit %s is run", param.key)
				ans, err := service.Judge(param.req)
				if err != nil {
					log.Error(err)
					return err
				}
				cache.GetJudgeResuCache().Update(param.key, *ans)
				return nil
			})
		} else {
			runtime.Gosched()
		}
	}
}

func (t *SubmitTask) AddTask(key string, req *pb.JudgeRequest) error {
	err := t.pool.Submit(func() error {
		log.Infof("submit %s is run", key)
		ans, err := service.Judge(req)
		if err != nil {
			log.Error(err)
			return err
		}
		cache.GetJudgeResuCache().Update(key, *ans)

		return nil
	})
	if err != nil {
		return err
	}
	log.Infof("add submittask uid: %v problemID %v\n", req.UserID, req.ProblemID)
	return nil
}
