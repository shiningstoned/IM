package ws

import (
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
	"im-demo/api/global"
	"im-demo/api/pkg/kafka"
	"im-demo/kitex_gen/message"
)

type Client struct {
	Uuid    string
	Conn    *websocket.Conn
	Send    chan []byte
	Storage chan message.SaveMsgRequset
}

func (c *Client) Read(uuid string) {
	defer func() {
		MyServer.Unregister <- c
		c.Conn.Close()
		global.Wg.Done()
	}()

	for {
		_, messages, err := c.Conn.ReadMessage()
		if err != nil {
			hlog.Fatalf("read message failed: %s", err.Error())
		}
		msg := message.Message{}
		json.Unmarshal(messages, &msg)
		msg.FromId = &uuid
		messages, _ = json.Marshal(msg)

		kafka.Send(messages)
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(1, message)
	}
}
