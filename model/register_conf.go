package model

type ConsulRegister struct {
	Register *ConsulRegisterConf `json:"register"`
}

type ConsulRegisterConf struct {
	ConsulAddress                  string   `json:"consul_address"`
	Service                        string   `json:"service"`
	Tag                            []string `json:"tag"`
	Port                           int      `json:"port"`
	DeregisterCriticalServiceAfter int      `json:"deregisterCriticalServiceAfter"`
	Interval                       int      `json:"interval"`
}
