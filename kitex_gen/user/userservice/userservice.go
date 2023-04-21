// Code generated by Kitex v0.5.1. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "im-demo/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Login":        kitex.NewMethodInfo(loginHandler, newUserServiceLoginArgs, newUserServiceLoginResult, false),
		"Register":     kitex.NewMethodInfo(registerHandler, newUserServiceRegisterArgs, newUserServiceRegisterResult, false),
		"AddFriend":    kitex.NewMethodInfo(addFriendHandler, newUserServiceAddFriendArgs, newUserServiceAddFriendResult, false),
		"DeleteFriend": kitex.NewMethodInfo(deleteFriendHandler, newUserServiceDeleteFriendArgs, newUserServiceDeleteFriendResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func addFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceAddFriendArgs)
	realResult := result.(*user.UserServiceAddFriendResult)
	success, err := handler.(user.UserService).AddFriend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceAddFriendArgs() interface{} {
	return user.NewUserServiceAddFriendArgs()
}

func newUserServiceAddFriendResult() interface{} {
	return user.NewUserServiceAddFriendResult()
}

func deleteFriendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceDeleteFriendArgs)
	realResult := result.(*user.UserServiceDeleteFriendResult)
	success, err := handler.(user.UserService).DeleteFriend(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceDeleteFriendArgs() interface{} {
	return user.NewUserServiceDeleteFriendArgs()
}

func newUserServiceDeleteFriendResult() interface{} {
	return user.NewUserServiceDeleteFriendResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *user.LoginRequest) (r *user.LoginResponse, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Register(ctx context.Context, req *user.RegisterRequest) (r *user.CommonResponse, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddFriend(ctx context.Context, req *user.AddFriendRequest) (r *user.CommonResponse, err error) {
	var _args user.UserServiceAddFriendArgs
	_args.Req = req
	var _result user.UserServiceAddFriendResult
	if err = p.c.Call(ctx, "AddFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteFriend(ctx context.Context, req *user.DeleteFriendRequest) (r *user.CommonResponse, err error) {
	var _args user.UserServiceDeleteFriendArgs
	_args.Req = req
	var _result user.UserServiceDeleteFriendResult
	if err = p.c.Call(ctx, "DeleteFriend", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
