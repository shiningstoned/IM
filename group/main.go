package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"im-demo/group/initialize"
	"im-demo/group/pkg"
	group "im-demo/kitex_gen/group/groupservice"
	"log"
)

func main() {
	initialize.InitConfig()
	r, info := initialize.InitRegistry()
	db := initialize.InitMysql()
	client := initialize.InitRedis()

	svr := group.NewServer(&GroupServiceImpl{
		pkg.NewMysqlManger(db),
		pkg.NewRedisManager(client),
	},
		server.WithServiceAddr(utils.NewNetAddr("tcp", "172.24.111.215:8881")),
		server.WithRegistry(r),
		server.WithRegistryInfo(&info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "group_srv"}),
	)

	err := svr.Run()

	if err != nil {
		log.Fatal(err)
	}
}
