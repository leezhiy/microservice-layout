// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"git.corp.doulaoban.com/backend/kratos-layout/internal/biz"
	"git.corp.doulaoban.com/backend/kratos-layout/internal/conf"
	"git.corp.doulaoban.com/backend/kratos-layout/internal/data"
	"git.corp.doulaoban.com/backend/kratos-layout/internal/server"
	"git.corp.doulaoban.com/backend/kratos-layout/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, nacosConfig *conf.NacosConfig, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	orderRepo := data.NewOrderRepo(dataData, logger)
	orderUseCase := biz.NewOrderUseCase(orderRepo, logger)
	orderService := service.NewOrderService(orderUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, orderService)
	registrar := server.NewRegistrar(nacosConfig, confServer)
	app := newApp(confServer, logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
