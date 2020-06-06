package main

import (
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	log "github.com/sirupsen/logrus"
	"goGRPC/pb"
	"goGRPC/service"
	"goGRPC/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
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

	log.WithFields(log.Fields{
		"port": *port,
	}).Info("start server....")

	if err := server.Serve(sock); err != nil {
		util.PanicIfError("fail to start grpc server", err)
	}
}
