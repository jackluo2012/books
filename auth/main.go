package main

import (
	"books/auth/handler"
	"books/auth/model"
	s "books/auth/proto/auth"
	"books/basic"
	"books/basic/common"
	"books/basic/config"
	z "books/plugins/zap"
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc/v2"
)

var (
	appName = "auth_srv"
	cfg     = &authCfg{}
	log     = z.GetLogger()
)

type authCfg struct {
	common.AppCfg
}

func main() {

	// 初始化配置、数据库信息
	initCfg()

	// 使用etcd 注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()

			return nil
		}),
	)
	//注册服务
	//auth.RegisterAuthHandler(service.Server(), new(handler.Service))
	s.RegisterAuthHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal("启动服务失败，错误: " + err.Error())
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
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

	return
}
