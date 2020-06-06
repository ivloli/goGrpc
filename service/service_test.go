package service

import (
	"context"
	"goGRPC/pb"
	"goGRPC/util"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestHelloService(t *testing.T) {
	addr := "localhost:8090"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	util.PanicIfError("fail to dial grpc server", err)
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 500)
	defer cancel()

	req := hello.HelloRequest{
		Name: "Tom Clay",
	}
	resp, err := client.SayHello(ctx, &req)
	if err != nil {
		log.Printf("fail to get data from server:%v", err)
		return
	}
	log.Printf("resp:%v", resp.Reply)
}
