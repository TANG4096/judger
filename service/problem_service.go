package service

import (
	"context"
	"fmt"
	"judger/pb"
	"judger/register_center"

	"github.com/sta-golang/go-lib-utils/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

func Judge(req *pb.JudgeRequest) (result *string, err error) {
	schema, err := register_center.GenerateAndRegisterEtcdResolver("127.0.0.1:2379", "JudgeService")
	if err != nil {
		log.Fatal("init etcd resolver err:", err.Error())
		return nil, err
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:///JudgeService", schema), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewJudgeServiceClient(conn)
	log.Infof("rpc client created")
	log.Infof("uid: %d code:%s", req.UserID, req.SourceCode)
	res, err := c.Judge(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	log.Infof("result:%v\n", res.Response)
	return &res.Response, nil
}
