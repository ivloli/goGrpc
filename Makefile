run_server:
	go run cmd/server/simple/server.go

run_client:
	go run cmd/client/simple/client.go

run_consul_server:
	go run cmd/server/consul/server.go

run_consul_client:
	go run cmd/client/consul/client.go

run_etcd_server:
	go run cmd/server/etcd/server.go

run_etcd_client:
	go run cmd/client/etcd/client.go