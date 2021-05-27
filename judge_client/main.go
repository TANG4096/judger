package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"judger/pb"
	"judger/register_center"
	"judger/util"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

var sum int64 = 0

func test() time.Duration {
	//连接etcd,得到名命名空间
	t1 := time.Now()
	schema, err := register_center.GenerateAndRegisterEtcdResolver("127.0.0.1:2379", "JudgeService")
	if err != nil {
		log.Fatal("init etcd resolver err:", err.Error())
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:///JudgeService", schema), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	//fmt.Println(1)
	//创建客户端存根对象
	c2 := pb.NewJudgeServiceClient(conn)
	req := &pb.JudgeRequest{
		ProblemID: 5,
		Type:      "C++",
		IsUpdate:  false,
	}
	path := util.GetPath() + "/temp_data/test/"
	sourceFile, err := ioutil.ReadFile(path + "main.cpp")
	if err != nil {
		fmt.Println(err)
	}
	req.SourceCode = sourceFile

	_, err = c2.Judge(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	//fmt.Printf("result:%v\n", res2.Response)
	defer conn.Close()
	d := time.Since((t1))
	//fmt.Printf("耗时：%v\n", d)

	atomic.AddInt64(&sum, int64(d))
	//wg.Done()
	return d
}

func main() {
	wg := sync.WaitGroup{}
	var n int64 = 100
	for i := 1; i <= int(n); i++ {
		time.Sleep(100 * time.Microsecond)
		test()
	}
	wg.Wait()
	fmt.Printf("\n请求次数： %d次\n总耗时：%v\n", 1000, time.Duration(sum))
	var an float64 = float64(sum)
	an = an / float64(n)
	fmt.Printf("平均耗时：%v\n", time.Duration(int64(an)))
}
