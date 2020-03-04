# User Service

This is the User service

Generated with

```
micro new books/user-srv --namespace=mu.micro.book --alias=user --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.srv.user
- Type: srv
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```
### docker run etc
```bash
docker kill etcd
docker rm etcd
export NODE1=0.0.0.0
docker run -d \
  -p 2379:2379 \
  -p 2380:2380 \
  --volume=${DATA_DIR}:/etcd-data \
  --name etcd quay.io/coreos/etcd:latest \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://${NODE1}:2380 \
  --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://${NODE1}:2379 \
  --initial-cluster node1=http://${NODE1}:2380
```

### 生成协议文件
```
protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto
```

### 运行
```
 go run main.go plugin.go --registry=etcd --registry_address=127.0.0.1:2379
```
### 测试
```
micro --registry=etcd call mu.micro.book.srv.user User.QueryUserByName '{"userName":"micro"}'
```
## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-srv
```

Build a docker image
```
make docker
```