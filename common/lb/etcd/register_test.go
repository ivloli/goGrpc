package etcd

import (
	"github/leel0330/grpcdemo/model"
	"log"
	"os"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	info := model.ServerNodeInfo{
		Name:           "greeter",
		Version:        "v1",
		Addr:           "127.0.0.1:8080",
		Weight:         1,
		LastUpdateTime: time.Now(),
	}

	r := EtcdRegister{
		EtcdAddrs:   []string{"127.0.0.1:2379"},
		DialTimeout: 3,
	}

	closeCh, err := r.Register(info, 10)
	if err != nil {
		t.Fatalf("Register to etcd failed." + err.Error())
	}

	infoEtcd, err := r.GetServiceInfo()
	if err != nil {
		t.Fatalf("Get from etcd failed.")
	}

	log.Printf("From etcd:%v\n", infoEtcd.Addr)
	time.Sleep(5 * time.Second)

	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("fail to get hostname:%v", err)
	} else {
		log.Printf("hostnameL:%v", hostname)
	}

	closeCh <- struct{}{}
}
