# IM
即时聊天demo微服务版
# 功能实现
* 注册
* 登录
* 添加好友
* 删除好友
* 创建群聊
* 加入群聊
* 单聊，群聊
# 后端框架
* kitex: cloudwego的微服务rpc框架
* hertz: cloudwego的微服务http框架
* gorm: orm框架
* mysql: 数据库(用于存储用户，群聊信息)
* redis: 数据库(用于存储好友关系，群聊关系)
* kafka: 消息队列
# 目录结构
``` 
.
├── api                                    api网关
│   ├── biz
│   │   ├── handler
│   │   │   ├── api
│   │   │   │   └── api_service.go         api实现
│   │   │   └── ping.go
│   │   ├── model
│   │   │   └── api
│   │   │       └── api.go
│   │   └── router
│   │       ├── api
│   │       │   ├── api.go
│   │       │   └── middleware.go
│   │       └── register.go
│   ├── global
│   │   └── global.go
│   ├── initialize                          服务发现
│   │   ├── register.go
│   │   └── rpc
│   │       ├── group.srv.go
│   │       ├── init.go
│   │       ├── message_srv.go
│   │       └── user_srv.go
│   ├── main.go
│   ├── middleware                           中间件 
│   │   ├── cors.go
│   │   └── jwt.go
│   ├── pkg
│   │   ├── constant                         常量
│   │   │   └── constant.go
│   │   └── kafka                            初始化kafka
│   │       ├── consumer.go
│   │       ├── init.go
│   │       └── producer.go
│   |
│   └── ws                                   websocket处理
│       ├── client.go
│       └── server.go
├── go.mod
|
├── group                                    group服务 创建组和加入组
│   |
│   ├── config
│   │   └── config.toml
│   ├── handler.go
│   ├── initialize
│   │   ├── config.go
│   │   ├── db.go
│   │   └── registry.go
│   |
│   ├── main.go
│   ├── pkg
│       ├── mysql.go
│       └── redis.go
|
├── idl
│   ├── api.thrift
│   ├── group.thrift
│   ├── message.thrift
│   └── user.thrift
|
├── kitex_gen
│  
├── message                                   message服务 存储消息
│   ├── config
│   │   └── config.toml
│   ├── handler.go
│   ├── initialize
│   │   ├── config.go
│   │   ├── db.go
│   │   └── register.go
│   ├── main.go
│   ├── pkg
│       └── mysql.go
│  
└── user                                       user服务 创建用户，添加好友
    ├── config
    │   └── config.toml
    ├── handler.go
    ├── initialize
    │   ├── config.go
    │   ├── db.go
    │   └── register.go
    ├
    ├── main.go
    ├── pkg
        ├── mysql.go
        └── redis.go
```
# 代码实现
## 升级websocket
``` 
uuid, _ := c.Get("uuid")
	err := global.Upgrader.Upgrade(c, func(conn *websocket.Conn) {
		client := &ws.Client{
			Uuid:    uuid.(string),
			Conn:    conn,
			Send:    make(chan []byte),
			Storage: make(chan message.SaveMsgRequset, 1),
		}
		global.Wg.Add(1)
		ws.MyServer.Register <- client
		go client.Read(uuid.(string))
		go client.Write()
		global.Wg.Wait()
	})
 ```
 ## client和server 
 ``` 
 type Client struct {
	Uuid    string
	Conn    *websocket.Conn
	Send    chan []byte
	Storage chan message.SaveMsgRequset
}

type Server struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Broadcast  chan []byte
}
```
## 单聊与群聊
```
// 单聊
conn, ok := s.Clients[msg.ToId]
				if ok {
					conn.Send <- messages
				}
        
//群聊 
groupMem, err := global.GroupClinet.GetGroupMem(context.Background(), &group.GetGroupMemRequest{GroupUuid: msg.ToId})
				if err != nil {
					hlog.Fatalf("get group members failed: %s", err.Error())
				}
				for _, mem := range groupMem.GroupMemUuid {
          conn, ok := s.Clients[mem]
					if ok {
						conn.Send <- messages
					}
        }
```
## 图片处理
```
name := "image/" + *msg.FromId + msg.ToId + time.Now().Format("2006-01-02 15:04:05") + ".jpg"
	encodedData := strings.Split(msg.Content, ",")[1]
	imageData, err := base64.StdEncoding.DecodeString(encodedData)

	file, err := os.Create(name)
	if err != nil {
		log.Fatal("File creation error:", err)
	}

	_, err = file.Write(imageData)
	if err != nil {
		log.Fatal("File write error:", err)
	}
	file.Close()
```
