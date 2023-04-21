package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
	user "im-demo/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	MysqlManager
	RedisManager
}

type RedisManager interface {
	AddFriend(myUuid, friendUuid string) (bool, error)
	DelFriend(myUuid, friendUuid string) (bool, error)
}

type MysqlManager interface {
	CreateUser(username, password string) error
	LoginCheck(username, password string) (bool, string, error)
	IsUserExist(uuid string) (bool, error)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	flag, uuid, err := s.MysqlManager.LoginCheck(req.Username, req.Password)
	if err != nil {
		klog.Error("login check error, err: ", err)
		return nil, status.Err(codes.Internal, "login error")
	}
	if !flag {
		klog.Info("wrong password")
		return nil, status.Err(codes.Internal, "login error")
	}
	return &user.LoginResponse{Userid: uuid}, nil
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.CommonResponse, err error) {
	// TODO: Your code here...
	err = s.MysqlManager.CreateUser(req.Username, req.Password)
	if err != nil {
		klog.Error("login check error, err:", err)
		return nil, status.Err(codes.Internal, "register error")
	}
	return &user.CommonResponse{Response: "register success"}, nil
}

// AddFriend implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddFriend(ctx context.Context, req *user.AddFriendRequest) (resp *user.CommonResponse, err error) {
	// TODO: Your code here...
	exist, err := s.MysqlManager.IsUserExist(req.FriendId)
	if err != nil {
		klog.Error("search user error, err:", err)
		return nil, status.Err(codes.Internal, "search user error")
	}
	if !exist {
		return &user.CommonResponse{Response: "user not exist"}, nil
	}
	flag, err := s.RedisManager.AddFriend(req.MyId, req.FriendId)
	if err != nil {
		klog.Error("add friend error, err:", err)
		return nil, status.Err(codes.Internal, "add friend failed")
	}
	if flag == false {
		return &user.CommonResponse{Response: "you are already friend!"}, nil
	}
	return &user.CommonResponse{Response: "add friend success"}, nil
}

// DeleteFriend implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteFriend(ctx context.Context, req *user.DeleteFriendRequest) (resp *user.CommonResponse, err error) {
	// TODO: Your code here...
	exist, err := s.MysqlManager.IsUserExist(req.FriendId)
	if err != nil {
		klog.Error("search user error, err:", err)
		return nil, status.Err(codes.Internal, "search user error")
	}
	if !exist {
		return &user.CommonResponse{Response: "user not exist"}, nil
	}
	flag, err := s.RedisManager.DelFriend(req.MyId, req.FriendId)
	if err != nil {
		klog.Error("delete friend error, err:", err)
		return nil, status.Err(codes.Internal, "delete friend error")
	}
	if flag == false {
		return &user.CommonResponse{Response: "you are not friend!"}, nil
	}
	return &user.CommonResponse{Response: "delete friend success"}, nil
}
