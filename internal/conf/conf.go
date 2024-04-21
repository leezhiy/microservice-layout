package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	nconfig "git.corp.doulaoban.com/backend/kratos-layout/pkg/nacos/config"
)

func NewNacosSource(conf *NacosConfig, appConfig *Server) config.Source {
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
	client, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	return nconfig.NewConfigSource(client,
		nconfig.WithGroup(conf.Group),
		nconfig.WithDataID(appConfig.Name),
		nconfig.WithFileExtension(conf.Extension),
	)

}
