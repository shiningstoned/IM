package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	message "im-demo/kitex_gen/message"
	"im-demo/message/pkg"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct {
	MysqlManager
}

type MysqlManager interface {
	SaveMsg(msg pkg.Message) error
}

// SaveMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SaveMessage(ctx context.Context, req *message.SaveMsgRequset) (resp *message.SaveMsgResponse, err error) {
	// TODO: Your code here...
	var message = pkg.Message{
		FromId:  req.FromId,
		ToId:    req.ToId,
		Content: req.Content,
	}
	err = s.MysqlManager.SaveMsg(message)
	if err != nil {
		klog.Fatalf("save message failed: %s", err.Error())
	}
	return nil, err
}
