package global

import (
	"github.com/hertz-contrib/websocket"
	"im-demo/kitex_gen/message/messageservice"
	"im-demo/kitex_gen/user/userservice"
	"sync"
)

var (
	UserClient    userservice.Client
	MessageClient messageservice.Client
	Upgrader      websocket.HertzUpgrader
	Wg            sync.WaitGroup
)
