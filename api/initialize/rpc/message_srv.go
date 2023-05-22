package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consul "github.com/kitex-contrib/registry-consul"
	message "im-demo/kitex_gen/message/messageservice"
)

func InitMessageService() message.Client {
	r, err := consul.NewConsulResolver("172.24.111.215:8500")
	if err != nil {
		hlog.Fatalf("new consul resolver failed: %s", err.Error())
	}

	cli, err := message.NewClient("message_srv",
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "message_srv"}),
	)
	if err != nil {
		hlog.Fatalf("new message client failed: %s", err.Error())
	}
	return cli
}
