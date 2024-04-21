package main

import (
	"github.com/go-kratos/kratos/v2" // 三方包
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/uuid"

	"git.corp.doulaoban.com/backend/kratos-layout/internal/conf"
	boot "git.corp.doulaoban.com/backend/kratos-layout/internal/init"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the application.
	Version string
	id      = uuid.New().String()
)

func newApp(conf *conf.Server, logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(conf.Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	// 加载本地配置文件
	localConfigClient, err := boot.Config().LoadLocalConfig()
	if err != nil {
		panic(err)
	}

	// 加载 Nacos 配置文件 bc is bootstrapConfig
	bc, err := localConfigClient.LoadRemoteConfig()
	if err != nil {
		panic(err)
	}

	// 初始化链路追踪
	if err := boot.Tracer(bc.Tracer, bc.Server); err != nil {
		panic(err)
	}

	// 初始化阿里云日志
	logger := boot.Logger(bc.Logger, bc.Server, id, Version)

	// 初始化 sentry
	if err := boot.Sentry(bc.Sentry, logger); err != nil {
		panic(err)
	}

	// 初始化应用
	app, cleanup, err := initApp(bc.Server, bc.Nacos.Discovery, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

}
