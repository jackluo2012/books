# User Service

This is the User service

Generated with

```
micro new books/user-web --namespace=mu.micro.book --alias=user --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.web.user
- Type: web
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

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-web
```

Build a docker image
```
make docker
```
### 运行 
```
 go run main.go --registry=etcd --registry_address=127.0.0.1:2379
```

### 运行进行测试
```
micro --registry=etcd --api_namespace=mu.micro.book.web  api --handler=web
```

