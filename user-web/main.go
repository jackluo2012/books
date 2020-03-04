package main

import (
	"books/basic"
	"books/basic/config"
	"books/user-web/handler"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

func main() {
	// 初始化配置
	basic.Init()

	//使用etcd 注册
	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)

	// initialise service 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				//初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// register call handler
	service.HandleFunc("/user/login", handler.Login)
	// 注册退出接口
	service.HandleFunc("/user/logout", handler.Logout)
	service.HandleFunc("/user/test", handler.TestSession)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
