package main

import (
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	log "github.com/sirupsen/logrus"
	"goGRPC/common/lb/etcd"
	"goGRPC/model"
	"goGRPC/pb"
	"goGRPC/service"
	"goGRPC/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

var (
	port = flag.Int("p", 8090, "grpc port")
)

func init() {
	flag.Parse()
	flag.VisitAll(func(flag *flag.Flag) {
		log.WithFields(log.Fields{
			"name":  flag.Name,
			"value": flag.Value,
		}).Info("flag params")
	})

}

func main() {
	r := etcd.EtcdRegister{
		EtcdAddrs:   []string{"127.0.0.1:2379"},
		DialTimeout: 3,
	}

	defer r.Stop()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	helloService := &service.HelloServiceImpl{}
	healthService := &service.HealthImpl{}

	hello.RegisterHelloServiceServer(server, helloService)
	grpc_health_v1.RegisterHealthServer(server, healthService)

	sock, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	util.PanicIfError("fail to listen port", err)

	srvName, version := "greeting", "v1"
	info := model.ServerNodeInfo{
		Name:           srvName,
		Version:        version,
		Addr:           fmt.Sprintf("127.0.0.1:%d", *port),
		Weight:         1,
		LastUpdateTime: time.Now(),
	}
	r.Register(info, 10)

	registerInfo, err := r.GetServiceInfo()
	if err == nil {
		log.Printf("register service ok:name=%v,addr=%v", registerInfo.Name, registerInfo.Addr)
	}

	log.WithFields(log.Fields{
		"port": *port,
	}).Info("start server....")

	reflection.Register(server)
	if err := server.Serve(sock); err != nil {
		util.PanicIfError("fail to start grpc server", err)
	}
}
