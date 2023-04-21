package pkg

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromId  string
	ToId    string
	Content string
}

type MessageManager struct {
	db *gorm.DB
}

func NewMessageManager(db *gorm.DB) *MessageManager {
	m := db.Migrator()
	if !m.HasTable(&Message{}) {
		err := m.CreateTable(&Message{})
		if err != nil {
			klog.Fatalf("create table failed: %s", err.Error())
		}
	}
	return &MessageManager{
		db: db,
	}
}

func (m *MessageManager) SaveMsg(msg Message) error {
	err := m.db.Create(&msg).Error
	if err != nil {
		return err
	}
	return nil
}
