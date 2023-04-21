package rpc

import (
	"im-demo/api/global"
)

func Init() {
	global.UserClient = InitUserService()
	global.MessageClient = InitMessageService()
}
