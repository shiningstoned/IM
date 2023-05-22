package pkg

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	GroupUuid string
	OwnerUuid string
	Name      string
	Notice    string
}

func (g *Group) BeforeCreate(_ *gorm.DB) error {
	if g.GroupUuid == "" {
		g.GroupUuid = uuid.New().String()
	}
	return nil
}

type MysqlManger struct {
	db *gorm.DB
}

func NewMysqlManger(db *gorm.DB) *MysqlManger {
	m := db.Migrator()
	if !m.HasTable(&Group{}) {
		err := m.CreateTable(&Group{})
		if err != nil {
			klog.Errorf("create table failed: %s", err.Error())
			return nil
		}
	}
	return &MysqlManger{
		db: db,
	}
}

func (m *MysqlManger) CreateGroup(OwnerUuid string) (string, error) {
	var group Group
	group.OwnerUuid = OwnerUuid
	err := m.db.Create(&group).Error
	if err != nil {
		return "", err
	}
	return group.GroupUuid, nil
}
