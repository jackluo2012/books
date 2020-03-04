package main

import (
	"books/auth/handler"
	"books/auth/model"
	auth "books/auth/proto/auth"
	"books/basic"
	"books/basic/config"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
)

func main() {

	// 初始化配置、数据库信息
	basic.Init()

	// 使用etcd 注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(micro.Action(
		func(c *cli.Context) {
			//初始化handler
			model.Init()
			//初始化 handler
			handler.Init()
		}), )
	//注册服务
	auth.RegisterAuthHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
