package main

import (
	"books/basic"
	"books/basic/config"
	"books/inventory-srv/handler"
	"books/inventory-srv/model"
	inventory "books/inventory-srv/proto/inventory"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.inventory"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(context *cli.Context) {
			//初始化模型层
			model.Init()
			//初始化handler
			handler.Init()
		}),
	)

	// Register Handler
	inventory.RegisterInventoryHandler(service.Server(), new(handler.Inventory))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
