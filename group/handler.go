package main

import (
	"context"
	group "im-demo/kitex_gen/group"
)

// GroupServiceImpl implements the last service interface defined in the IDL.
type GroupServiceImpl struct {
	MysqlManager
	RedisManager
}

type MysqlManager interface {
	CreateGroup(OwnerUuid string) (string, error)
}

type RedisManager interface {
	JoinGroup(groupUuid, myUuid string) error
	GetGroupMem(groupUuid string) ([]string, error)
	IsGroupMem(groupUuid, myUuid string) bool
}

// CreateGroup implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) CreateGroup(ctx context.Context, req *group.CreateGroupRequest) (resp *group.CommonResponse, err error) {
	// TODO: Your code here...
	groupUuid, err := s.MysqlManager.CreateGroup(req.OwnerUuid)
	if err != nil {
		return nil, err
	}
	err = s.RedisManager.JoinGroup(groupUuid, req.OwnerUuid)
	if err != nil {
		return nil, err
	}
	return &group.CommonResponse{Message: "create group success"}, nil
}

// JoinGroup implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) JoinGroup(ctx context.Context, req *group.JoinGroupRequest) (resp *group.CommonResponse, err error) {
	// TODO: Your code here...
	flag := s.RedisManager.IsGroupMem(req.GroupUuid, req.MyUuid)
	if flag {
		return &group.CommonResponse{Message: "you are the member"}, nil
	}
	err = s.RedisManager.JoinGroup(req.GroupUuid, req.MyUuid)
	if err != nil {
		return nil, err
	}
	return &group.CommonResponse{Message: "join group success"}, nil
}

// GetGroupMem implements the GroupServiceImpl interface.
func (s *GroupServiceImpl) GetGroupMem(ctx context.Context, req *group.GetGroupMemRequest) (resp *group.GetGroupMemResponse, err error) {
	// TODO: Your code here...
	mem, err := s.RedisManager.GetGroupMem(req.GroupUuid)
	if err != nil {
		return nil, err
	}
	return &group.GetGroupMemResponse{GroupMemUuid: mem}, nil
}
