package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consul "github.com/kitex-contrib/registry-consul"
	group "im-demo/kitex_gen/group/groupservice"
)

func InitGroupService() group.Client {
	r, err := consul.NewConsulResolver("172.24.111.215:8500")
	if err != nil {
		hlog.Fatalf("new consul resolver failed: %s", err.Error())
	}

	client, err := group.NewClient("group_srv",
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "group_srv"}))
	if err != nil {
		hlog.Fatalf("new group service client failed: %s", err.Error())
	}
	return client
}
