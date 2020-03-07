package main

import (
	"books/basic"
	"books/basic/common"
	"books/basic/config"
	"books/orders-web/handler"
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/config/source/grpc/v2"
	"net/http"
	"github.com/micro/go-micro/util/log"
)


var (
	appName = "orders_web"
	cfg     = &appCfg{}
)

type appCfg struct {
	common.AppCfg
}

func main() {

	// 初始化配置
	initCfg()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 创建新服务
	// 创建新服务
	service := web.NewService(
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)
	log.Log("服务器的端口是: ",cfg.Addr())
	// 初始化服务
	if err := service.Init(
		web.BeforeStart(func() error {
			return nil
		}),
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
		web.AfterStart(func() error {
			return nil
		}),
	); err != nil {
		log.Fatal(err)
	}
	// 新建订单接口
	authHandler := http.HandlerFunc(handler.New)
	service.Handle("/orders/new", handler.AuthWrapper(authHandler))
	service.HandleFunc("/orders/hello", handler.Hello)

	// 运行服务
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
