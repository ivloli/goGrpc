package consul

import (
	"github.com/hashicorp/consul/api"
	"log"
	"testing"
)

func TestConsulClient(t *testing.T) {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Printf("fail to init consul client:%v", err)
		return
	}

	serviceName := "helloService"
	services, metaInfo, err := client.Health().ServiceMultipleTags(serviceName,
		[]string{"hello"}, true, &api.QueryOptions{
			WaitIndex: 0,
		})
	if err != nil {
		log.Printf("fail to get service:%v,%v", serviceName, err)
		return
	}
	log.Printf("last index:%v", metaInfo.LastIndex)
	for _, service := range services{
		log.Printf("%v-->%v:%v", service.Service.Service, service.Service.Address, service.Service.Port)
	}
}
