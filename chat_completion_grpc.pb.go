// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: chat_completion.proto

package chat_completion

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChatService_Ask_FullMethodName                 = "/chat_completion.ChatService/ask"
	ChatService_Chat_FullMethodName                = "/chat_completion.ChatService/chat"
	ChatService_WriteArticleByTitle_FullMethodName = "/chat_completion.ChatService/write_article_by_title"
	ChatService_TranscribeJudge_FullMethodName     = "/chat_completion.ChatService/transcribe_judge"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Ask(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ChatMessage, error)
	Chat(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatClient, error)
	WriteArticleByTitle(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (ChatService_WriteArticleByTitleClient, error)
	TranscribeJudge(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ChatMessage, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Ask(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ChatMessage, error) {
	out := new(ChatMessage)
	err := c.cc.Invoke(ctx, ChatService_Ask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_Chat_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChatClient{stream}
	return x, nil
}

type ChatService_ChatClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatServiceChatClient struct {
	grpc.ClientStream
}

func (x *chatServiceChatClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceChatClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) WriteArticleByTitle(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (ChatService_WriteArticleByTitleClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_WriteArticleByTitle_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceWriteArticleByTitleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_WriteArticleByTitleClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatServiceWriteArticleByTitleClient struct {
	grpc.ClientStream
}

func (x *chatServiceWriteArticleByTitleClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) TranscribeJudge(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*ChatMessage, error) {
	out := new(ChatMessage)
	err := c.cc.Invoke(ctx, ChatService_TranscribeJudge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Ask(context.Context, *ChatMessage) (*ChatMessage, error)
	Chat(ChatService_ChatServer) error
	WriteArticleByTitle(*ChatMessage, ChatService_WriteArticleByTitleServer) error
	TranscribeJudge(context.Context, *ChatMessage) (*ChatMessage, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Ask(context.Context, *ChatMessage) (*ChatMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ask not implemented")
}
func (UnimplementedChatServiceServer) Chat(ChatService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatServiceServer) WriteArticleByTitle(*ChatMessage, ChatService_WriteArticleByTitleServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteArticleByTitle not implemented")
}
func (UnimplementedChatServiceServer) TranscribeJudge(context.Context, *ChatMessage) (*ChatMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranscribeJudge not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Ask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Ask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Ask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Ask(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Chat(&chatServiceChatServer{stream})
}

type ChatService_ChatServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatServiceChatServer struct {
	grpc.ServerStream
}

func (x *chatServiceChatServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceChatServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_WriteArticleByTitle_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChatMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).WriteArticleByTitle(m, &chatServiceWriteArticleByTitleServer{stream})
}

type ChatService_WriteArticleByTitleServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type chatServiceWriteArticleByTitleServer struct {
	grpc.ServerStream
}

func (x *chatServiceWriteArticleByTitleServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_TranscribeJudge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).TranscribeJudge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_TranscribeJudge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).TranscribeJudge(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat_completion.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ask",
			Handler:    _ChatService_Ask_Handler,
		},
		{
			MethodName: "transcribe_judge",
			Handler:    _ChatService_TranscribeJudge_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "chat",
			Handler:       _ChatService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "write_article_by_title",
			Handler:       _ChatService_WriteArticleByTitle_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat_completion.proto",
}
