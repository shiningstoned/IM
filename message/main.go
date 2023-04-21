package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	message "im-demo/kitex_gen/message/messageservice"
	"im-demo/message/initialize"
	"im-demo/message/pkg"
	"log"
)

func main() {
	r, info := initialize.InitRegistry()
	initialize.InitConfig()
	db := initialize.InitMysql()

	svr := message.NewServer(&MessageServiceImpl{pkg.NewMessageManager(db)},
		server.WithServiceAddr(utils.NewNetAddr("tcp", "172.21.172.211:8882")),
		server.WithRegistry(r),
		server.WithRegistryInfo(&info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "message_srv"}),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
