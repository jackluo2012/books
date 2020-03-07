package main

import (
	"books/basic"
	"books/basic/common"
	"books/basic/config"
	"books/payment-srv/handler"
	"books/payment-srv/model"
	p "books/payment-srv/proto/payment"

	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc/v2"
)

var (
	appName = "payment_srv"
	cfg     = &appCfg{}
)

type appCfg struct {
	common.AppCfg
}

func main() {

	// 初始化配置、数据库等信息
	initCfg()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Version(cfg.Version),
		micro.Registry(micReg),
		micro.Address(cfg.Addr()),
	)

	// Initialise service
	service.Init(
		micro.Action(func(*cli.Context) error {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
			return nil
		}),
	)

	p.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// 启动服务
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
