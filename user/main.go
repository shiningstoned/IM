package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	user "im-demo/kitex_gen/user/userservice"
	"im-demo/user/initialize"
	"im-demo/user/pkg"
	"log"
)

func main() {
	initialize.InitConfig()
	r, info := initialize.InitRegistry()
	db := initialize.InitMysql()
	client := initialize.InitRedis()

	svr := user.NewServer(&UserServiceImpl{
		pkg.NewUserManager(db),
		pkg.NewRUserManager(client),
	},
		server.WithServiceAddr(utils.NewNetAddr("tcp", "172.24.111.215:8880")),
		server.WithRegistry(r),
		server.WithRegistryInfo(&info),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user_srv"}),
	)

	err := svr.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}
