package consul

import (
	"fmt"
	"goGRPC/model"
	"time"

	"github.com/hashicorp/consul/api"
)

// NewConsulRegister create a new consul register
func NewConsulRegister(conf *model.ConsulRegisterConf) *ConsulRegister {
	return &ConsulRegister{
		ConsulAddress:                        conf.ConsulAddress,
		Service:                        conf.Service,
		Tag:                            conf.Tag,
		Port:                           conf.Port,
		DeregisterCriticalServiceAfter: time.Duration(conf.DeregisterCriticalServiceAfter) * time.Second,
		Interval:                       time.Duration(conf.Interval) * time.Second,
	}
}

// ConsulRegister consul service register
type ConsulRegister struct {
	ConsulAddress                        string
	Service                        string
	Tag                            []string
	Port                           int
	DeregisterCriticalServiceAfter time.Duration
	Interval                       time.Duration
}

// Register register service
func (r *ConsulRegister) Register() error {
	config := api.DefaultConfig()
	config.Address = r.ConsulAddress
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}
	agent := client.Agent()

	// serviceIP := util.LocalIP()
	serviceIP := "127.0.0.1"

	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", r.Service, serviceIP, r.Port), // 服务节点的名称
		Name:    r.Service,    // 服务名称
		Tags:    r.Tag,                                          // tag，可以为空
		Port:    r.Port,                                         // 服务端口
		Address: serviceIP,                                             // 服务 IP
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval:                       r.Interval.String(),                            // 健康检查间隔
			GRPC:                           fmt.Sprintf("%v:%v/%v", serviceIP, r.Port, r.Service), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),      // 注销时间，相当于过期时间
		},
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}

	return nil
}
