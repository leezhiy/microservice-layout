package init

import (
	"flag"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"

	"git.corp.doulaoban.com/backend/kratos-layout/internal/conf"
)

type Configures struct {
	Config *conf.Bootstrap
}

var ConfPath string

func init() {
	flag.StringVar(&ConfPath, "conf", "configs", "config path, eg: -conf config.yaml")
	flag.Parse()
}

func (c *Configures) LoadLocalConfig() (*Configures, error) {
	// 加载配置文件
	cfg := config.New(
		config.WithSource(
			file.NewSource(ConfPath),
		))

	// 配置加载失败则直接返回
	if err := cfg.Load(); err != nil {
		return nil, err
	}

	if err := cfg.Scan(c.Config); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Configures) LoadRemoteConfig() (*conf.Bootstrap, error) {
	// 加载 Nacos 配置文件
	cfg := config.New(
		config.WithSource(
			// 方便远程覆盖本地配置 进行合并
			file.NewSource(ConfPath),
			conf.NewNacosSource(c.Config.Nacos.Configuration, c.Config.Server),
		),
	)

	if err := cfg.Load(); err != nil {
		return nil, err
	}

	if err := cfg.Scan(c.Config); err != nil {
		return nil, err
	}

	return c.Config, nil
}
