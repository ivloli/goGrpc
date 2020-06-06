# goGRPC

## 1. 功能实现
+ 实验性的Hello服务
+ consul服务发现与服务注册(可选)
+ etcd服务发现与服务注册(可选)

## 2. 如何使用

### 1. build proto file

> protoc -I proto/ proto/hello.proto --go_out=plugins=grpc:pb

### 2. 运行server

> make run_server

### 3. 运行client

> make run_client

### 4. 依赖服务运行

#### 4.1 etcd

> docker run --rm -itd -p 2379:2379 -p 2380:2380 -p 4001:4001 -p 7001:7001 \
  -v ~/data/db/etcd:/data --name vetcd quay.io/coreos/etcd

#### 4.1 consul

> consul agent -dev