package main

import (
	"log"
	"net"
	"time"

	"judger/impl"
	"judger/pb"
	"judger/register_center"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//构造注册中心对象
	etcdRegister := register_center.NewEtcdRegisterImpl("127.0.0.1:2379")
	//开始注册
	go func() {
		for {
			etcdRegister.Register(register_center.ServiceDescInfo{ServiceName: "JudgeService",
				Host: "127.0.0.1", Port: 8090, IntervalTime: time.Duration(10)})

			time.Sleep(time.Second * 5)
		}
	}()
	//创建一个grpc服务器对象
	gRpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(gRpcServer, &impl.HelloServiceServer{})
	pb.RegisterJudgeServiceServer(gRpcServer, &impl.JudgeServiceServer{})
	//开启服务端
	gRpcServer.Serve(lis)
}
