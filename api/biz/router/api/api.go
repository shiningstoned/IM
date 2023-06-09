// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "im-demo/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/wschat", append(_wschatMw(), api.WsChat)...)
	{
		_group := root.Group("/group", _groupMw()...)
		_group.POST("/create", append(_creategroupMw(), api.CreateGroup)...)
		_group.POST("/join", append(_joingroupMw(), api.JoinGroup)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.POST("/addfriend", append(_addfriendMw(), api.AddFriend)...)
		_user.POST("/delfriend", append(_delfriendMw(), api.DelFriend)...)
		_user.POST("/login", append(_loginMw(), api.Login)...)
		_user.POST("/register", append(_registerMw(), api.Register)...)
	}
}
