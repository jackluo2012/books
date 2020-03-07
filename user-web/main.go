package main

import (
	"books/basic"
	"books/basic/common"
	"books/basic/config"
	"books/user-web/handler"
	"fmt"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/config/source/grpc/v2"
)

var (
	appName = "user_web"
	cfg     = &userCfg{}
)

type userCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()

	//使用etcd 注册
	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)

	// initialise service 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				//初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal("这里报错了吗？？",err)
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
