package task

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

const DefaultCleanIntervalTime int = int(5 * time.Second)

type f func() error
type sig struct{}

type Worker struct {
	// pool who owns this worker.
	pool *TaskPool

	// task is a job should be done.
	task chan f

	// recycleTime will be update when putting a worker back into queue.
	recycleTime time.Time
}

type TaskPool struct {
	capacity int32
	// running is the number of the currently running goroutines.
	running int32
	// expiryDuration set the expired time (second) of every worker.
	expiryDuration time.Duration
	// workers is a slice that store the available workers.
	workers []*Worker
	// release is used to notice the pool to closed itself.
	release chan sig
	// lock for synchronous operation.
	lock sync.RWMutex
}

// NewPool generates a instance of ants pool
func NewPool(size int) (*TaskPool, error) {
	return NewTimingPool(size, DefaultCleanIntervalTime)
}

// NewTimingPool generates a instance of ants pool with a custom timed task
func NewTimingPool(size, expiry int) (*TaskPool, error) {
	if size <= 0 {
		return nil, errors.New("InvalidPoolSize")
	}
	if expiry <= 0 {
		return nil, errors.New("ErrInvalidPoolExpiry")
	}
	p := &TaskPool{
		capacity:       int32(size),
		release:        make(chan sig, 1),
		expiryDuration: time.Duration(expiry) * time.Second,
	}
	return p, nil
}

func (t *TaskPool) Running() int32 {

	return t.running
}

func (t *TaskPool) incRunning() {
	atomic.AddInt32(&t.running, 1)
}

func (t *TaskPool) decRunning() {
	atomic.AddInt32(&t.running, -1)
}

// Submit submit a task to pool
func (p *TaskPool) Submit(task f) error {
	if len(p.release) > 0 {
		return errors.New("ErrPoolClosed")
	}
	w := p.getWorker()
	w.task <- task
	return nil
}

func (p *TaskPool) getWorker() *Worker {
	var w *Worker
	// 标志变量，判断当前正在运行的worker数量是否已到达Pool的容量上限
	waiting := false
	// 加锁，检测队列中是否有可用worker，并进行相应操作
	p.lock.Lock()
	idleWorkers := p.workers
	n := len(idleWorkers) - 1
	// 当前队列中无可用worker
	if n < 0 {
		// 判断运行worker数目已达到该Pool的容量上限，置等待标志
		waiting = p.Running() >= p.capacity
	} else {
		// 当前队列有可用worker，从队列尾部取出一个使用
		w = idleWorkers[n]
		idleWorkers[n] = nil
		p.workers = idleWorkers[:n]
	}
	// 检测完成，解锁
	p.lock.Unlock()
	// Pool容量已满，新请求等待
	if waiting {
		// 利用锁阻塞等待直到有空闲worker
		for {
			p.lock.Lock()
			idleWorkers = p.workers
			l := len(idleWorkers) - 1
			if l < 0 {
				p.lock.Unlock()
				continue
			}
			w = idleWorkers[l]
			idleWorkers[l] = nil
			p.workers = idleWorkers[:l]
			p.lock.Unlock()
			break
		}
		// 当前无空闲worker但是Pool还没有满，
		// 则可以直接新开一个worker执行任务
	} else if w == nil {
		w = &Worker{
			pool: p,
			task: make(chan f, 1),
		}
		w.run()
		// 运行worker数加一
		p.incRunning()
	}
	return w
}

func (w *Worker) run() {
	go func() {
		// 循环监听任务列表，一旦有任务立马取出运行
		for f := range w.task {
			if f == nil {
				// 退出goroutine，运行worker数减一
				w.pool.decRunning()
				return
			}
			f()
			// worker回收复用
			w.pool.putWorker(w)
		}
	}()
}

func (p *TaskPool) putWorker(worker *Worker) {
	// 写入回收时间，亦即该worker的最后一次结束运行的时间
	worker.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, worker)
	p.lock.Unlock()
}
