### 搭建 consul
```
docker run -d --net=host -e 'CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt": true}'  --name consul_server  consul agent -server -bind=192.168.1.6 -bootstrap-expect=1  -node=node1 -client 0.0.0.0 -ui
```
### 利用proto 生成代码
```
protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto
```

### 创建 srv
```
micro new --namespace=mu.micro.book --type=srv --alias=user books/user-srv
```

### 创建web
```
micro new --namespace=mu.micro.book --type=web --alias=user books/user-web
```

```
$ micro --registry=etcd --api_namespace=mu.micro.book.web  api --handler=web

$ cd user-srv
$ go run main.go plugin.go

$ cd user-web  
$ go run main.go plugin.go
 

$ cd payment-srv
$ go run main.go plugin.go

$ cd payment-web
$ go run main.go plugin.go

$ cd orders-web
$ go run main.go plugin.go

$ cd orders-srv
$ go run main.go plugin.go

$ cd inventory-srv
$ go run main.go plugin.go

$ cd auth
$ go run main.go plugin.go

```