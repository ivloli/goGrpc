package main

import (
	"context"
	"fmt"
	"goGRPC/common/lb/etcd"
	"goGRPC/pb"
	"goGRPC/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

func main() {

	srvName, version := "greeting", "v1"

	r := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"}, srvName, version, int64(10))
	resolver.Register(r)

	target := fmt.Sprintf("%s:///%s", r.Scheme(), srvName)

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	conn, err := grpc.DialContext(ctx, target,
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
