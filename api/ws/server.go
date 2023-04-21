package ws

import (
	"encoding/json"
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
			var m = message.SaveMsgRequset{
				FromId:  *msg.FromId,
				ToId:    msg.ToId,
				Content: msg.Content,
			}
			s.Clients[*msg.FromId].Storage <- m
			conn, ok := s.Clients[msg.ToId]
			if ok {
				conn.Send <- messages
			}
		}
	}

}
