// Code generated by Kitex v0.5.2. DO NOT EDIT.

package groupservice

import (
	server "github.com/cloudwego/kitex/server"
	group "im-demo/kitex_gen/group"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler group.GroupService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}