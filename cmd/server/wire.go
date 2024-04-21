//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"git.corp.doulaoban.com/backend/kratos-layout/internal/biz"
	"git.corp.doulaoban.com/backend/kratos-layout/internal/conf"
	"git.corp.doulaoban.com/backend/kratos-layout/internal/data"
	"git.corp.doulaoban.com/backend/kratos-layout/internal/server"
	"git.corp.doulaoban.com/backend/kratos-layout/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.NacosConfig, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
