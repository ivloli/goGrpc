package main

import (
	"context"
	"fmt"
	"goGRPC/common/lb/consul"
	"goGRPC/pb"
	"goGRPC/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"time"
)

func main() {
	consul.Init([]string{"hello"})

	// localIP := util.LocalIP()
	localIP := "127.0.0.1"

	target := fmt.Sprintf("%v://%v:%v/%v", "consul", localIP, 8500, "helloService")

	ctxDial, cancelDial := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelDial()
	conn, err := grpc.DialContext(ctxDial, target,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)

	util.PanicIfError("fail to dial grpc server", err)
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	req := hello.HelloRequest{
		Name: "Tom Clay",
	}
	resp, err := client.SayHello(ctx, &req)
	util.PanicIfError("fail to call sayHello", err)
	log.Printf("resp:%v", resp.Reply)
}
