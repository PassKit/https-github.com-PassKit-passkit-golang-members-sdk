// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: io/core/a_rpc_messages.proto

package io

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesClient interface {
	GetMessage(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Message, error)
	CreateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Id, error)
	UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteMessage(ctx context.Context, in *Id, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}

type messagesClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesClient(cc grpc.ClientConnInterface) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) GetMessage(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/io.Messages/getMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) CreateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/io.Messages/createMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) UpdateMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/io.Messages/updateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) DeleteMessage(ctx context.Context, in *Id, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/io.Messages/deleteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/io.Messages/sendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServer is the server API for Messages service.
// All implementations should embed UnimplementedMessagesServer
// for forward compatibility
type MessagesServer interface {
	GetMessage(context.Context, *Id) (*Message, error)
	CreateMessage(context.Context, *Message) (*Id, error)
	UpdateMessage(context.Context, *Message) (*emptypb.Empty, error)
	DeleteMessage(context.Context, *Id) (*emptypb.Empty, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
}

// UnimplementedMessagesServer should be embedded to have forward compatible implementations.
type UnimplementedMessagesServer struct {
}

func (UnimplementedMessagesServer) GetMessage(context.Context, *Id) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessagesServer) CreateMessage(context.Context, *Message) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedMessagesServer) UpdateMessage(context.Context, *Message) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessagesServer) DeleteMessage(context.Context, *Id) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedMessagesServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

// UnsafeMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServer will
// result in compilation errors.
type UnsafeMessagesServer interface {
	mustEmbedUnimplementedMessagesServer()
}

func RegisterMessagesServer(s grpc.ServiceRegistrar, srv MessagesServer) {
	s.RegisterService(&Messages_ServiceDesc, srv)
}

func _Messages_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Messages/getMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).GetMessage(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Messages/createMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).CreateMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Messages/updateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).UpdateMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Messages/deleteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).DeleteMessage(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.Messages/sendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Messages_ServiceDesc is the grpc.ServiceDesc for Messages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMessage",
			Handler:    _Messages_GetMessage_Handler,
		},
		{
			MethodName: "createMessage",
			Handler:    _Messages_CreateMessage_Handler,
		},
		{
			MethodName: "updateMessage",
			Handler:    _Messages_UpdateMessage_Handler,
		},
		{
			MethodName: "deleteMessage",
			Handler:    _Messages_DeleteMessage_Handler,
		},
		{
			MethodName: "sendMessage",
			Handler:    _Messages_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "io/core/a_rpc_messages.proto",
}
