package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consul "github.com/kitex-contrib/registry-consul"
	user "im-demo/kitex_gen/user/userservice"
)

func InitUserService() user.Client {
	resolver, err := consul.NewConsulResolver("172.21.172.211:8500")
	if err != nil {
		hlog.Fatalf("new consul resolver failed: %s", err.Error())
	}

	Client, err := user.NewClient(
		"user_srv",
		client.WithResolver(resolver),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user_srv"}),
	)
	if err != nil {
		klog.Fatalf("new user service client failed: %s", err.Error())
	}
	return Client
}
