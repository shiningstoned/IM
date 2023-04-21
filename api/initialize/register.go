package initialize

import (
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
)

func InitRegistry() (registry.Registry, registry.Info) {
	cfg := api.DefaultConfig()
	cfg.Address = "172.21.172.211:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}
	r := consul.NewConsulRegister(client,
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "15s",
		}),
	)
	info := registry.Info{
		Weight:      10,
		ServiceName: "api",
		Addr:        utils.NewNetAddr("tcp", "172.21.172.211:8881"),
	}
	return r, info
}
