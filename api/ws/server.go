package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"im-demo/api/global"
	"im-demo/api/pkg/constant"
	"im-demo/kitex_gen/group"
	"im-demo/kitex_gen/message"
	"log"
)

var MyServer *Server

type Server struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string]*Client
	Broadcast  chan []byte
}

func NewServer() *Server {
	return &Server{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan []byte),
	}
}

func ConsumeKafkaMag(data []byte) {
	MyServer.Broadcast <- data
}

func (s *Server) Start() {
	for {
		select {
		case conn := <-s.Register:
			s.Clients[conn.Uuid] = conn
			from := "system"
			content := "hello client"
			msg := message.Message{
				FromId:  &from,
				ToId:    conn.Uuid,
				Content: content,
				Type:    constant.TYPE_SYSTEM,
			}
			message, _ := json.Marshal(msg)
			conn.Send <- message
		case conn := <-s.Unregister:
			log.Println("user unregister")
			if _, ok := s.Clients[conn.Uuid]; ok {
				close(conn.Send)
				delete(s.Clients, conn.Uuid)
			}
		case messages := <-s.Broadcast:
			msg := message.Message{}
			json.Unmarshal(messages, &msg)
			if msg.Type == constant.TYPE_USER {
				_, err := global.MessageClient.SaveMessage(context.Background(), &message.SaveMsgRequset{
					FromId:  *msg.FromId,
					ToId:    msg.ToId,
					Content: msg.Content,
					Type:    constant.TYPE_USER,
				})
				if err != nil {
					hlog.Errorf("save message failed: %s", err.Error())
				}
				conn, ok := s.Clients[msg.ToId]
				if ok {
					conn.Send <- messages
				}
			} else if msg.Type == constant.TYPE_GROUP {
				groupMem, err := global.GroupClinet.GetGroupMem(context.Background(), &group.GetGroupMemRequest{GroupUuid: msg.ToId})
				if err != nil {
					hlog.Fatalf("get group members failed: %s", err.Error())
				}
				fmt.Println(msg.ToId, groupMem)
				for _, mem := range groupMem.GroupMemUuid {
					if mem == *msg.FromId {
						global.MessageClient.SaveMessage(context.Background(), &message.SaveMsgRequset{
							FromId:  msg.ToId,
							ToId:    *msg.FromId,
							Content: msg.Content,
							Type:    constant.TYPE_GROUP,
						})
						continue
					}
					global.MessageClient.SaveMessage(context.Background(), &message.SaveMsgRequset{
						FromId:  msg.ToId,
						ToId:    mem,
						Content: msg.Content,
						Type:    constant.TYPE_GROUP,
					})
					conn, ok := s.Clients[mem]
					if ok {
						conn.Send <- messages
					}
				}
			}
		}
	}

}
