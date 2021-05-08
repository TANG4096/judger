package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"judger/pb"
	"judger/register_center"
	"judger/util"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

func main() {
	//连接etcd,得到名命名空间
	schema, err := register_center.GenerateAndRegisterEtcdResolver("127.0.0.1:2379", "JudgeService")
	if err != nil {
		log.Fatal("init etcd resolver err:", err.Error())
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:///JudgeService", schema), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(1)
	//创建客户端存根对象
	c2 := pb.NewJudgeServiceClient(conn)
	req := &pb.JudgeRequest{
		ProblemID: 3,
		Type:      "C++",
		IsUpdate:  false,
	}
	path := util.GetPath() + "/temp_data/test/"
	sourceFile, err := ioutil.ReadFile(path + "a.cpp")
	if err != nil {
		fmt.Println(err)
	}
	req.SourceCode = sourceFile
	res2, err := c2.Judge(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("result:%v\n", res2.Response)
	defer conn.Close()
}
