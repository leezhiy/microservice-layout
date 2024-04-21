package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"git.corp.doulaoban.com/backend/kratos-layout/internal/conf"
	knacos "git.corp.doulaoban.com/backend/kratos-layout/pkg/nacos/registry"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.NacosConfig, appConfig *conf.Server) registry.Registrar {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Ip, conf.Port, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(appConfig.Environment),
		constant.WithUsername(conf.Username),
		constant.WithPassword(conf.Password),
		constant.WithTimeoutMs(conf.TimeoutMs),
	)

	// create config client
	client, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	var group = constant.DEFAULT_GROUP
	if conf.Group != "" {
		group = conf.Group
	}

	return knacos.New(client, knacos.WithGroup(group))
}
