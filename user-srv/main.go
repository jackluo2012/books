package main

import (
	"books/basic"
	"books/basic/common"
	"books/basic/config"
	"books/user-srv/handler"
	"books/user-srv/model"
	user "books/user-srv/proto/user"
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/micro/go-plugins/config/source/grpc/v2"
)

/**
1.新建service
2.初始化
3.注册服务
4.启动服务
*/

var (
	appName = "user_srv"
	cfg     = &userCfg{}
)

type userCfg struct {
	common.AppCfg
}

func main() {

	//初始化配置、数据库等信息
	initCfg()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	log.Log("micReg=",micReg)
	log.Log("cfg.Name",cfg.Name)
	// New Service 新建service
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Version(cfg.Version),
		micro.Registry(micReg),
		micro.Address(cfg.Addr()),
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
		micro.Action(func(context *cli.Context) error {
			// 初始化模型层
			model.Init()
			//初始化handler
			handler.Init()
			return nil
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
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
	etcd.Auth(etcdCfg.User,etcdCfg.Pass)

	log.Log("ops.Addrs =",ops.Addrs )
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Infof("[initCfg] 配置，cfg：%v", cfg)

	return
}
