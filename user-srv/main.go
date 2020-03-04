package main

import (
	"books/basic"
	"books/basic/config"
	"books/user-srv/handler"
	"books/user-srv/model"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"time"

	user "books/user-srv/proto/user"
)

/**
1.新建service
2.初始化
3.注册服务
4.启动服务
*/
func main() {

	//初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// New Service 新建service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service 初始化
	service.Init(
		micro.BeforeStart(func() error {
			log.Log(" 启动前的日志");
			return nil
		}),
		micro.AfterStart(func() error {
			log.Log(" 启动后的日志");
			return nil
		}),
		micro.Action(func(context *cli.Context) {
			// 初始化模型层
			model.Init()
			//初始化handler
			handler.Init()
		}),
	)

	// Register Handler 注册服务
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
