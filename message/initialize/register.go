package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
)

func InitRegistry() (registry.Registry, registry.Info) {
	r, err := consul.NewConsulRegister("172.24.111.215:8500",
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "15s",
		}),
	)
	if err != nil {
		klog.Fatalf("new consul register failed: %s", err.Error())
	}

	info := registry.Info{
		ServiceName: "message_srv",
		Addr:        utils.NewNetAddr("tcp", "172.24.111.215:8882"),
		Weight:      10,
	}
	return r, info
}
