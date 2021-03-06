// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

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

// HelloWorldClient is the client API for HelloWorld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloWorldClient interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	//server-side streaming
	GreetInManyLanguages(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (HelloWorld_GreetInManyLanguagesClient, error)
	//client-side streaming
	GreetManyPeople(ctx context.Context, opts ...grpc.CallOption) (HelloWorld_GreetManyPeopleClient, error)
	//bidirectional streaming
	Chat(ctx context.Context, opts ...grpc.CallOption) (HelloWorld_ChatClient, error)
}

type helloWorldClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloWorldClient(cc grpc.ClientConnInterface) HelloWorldClient {
	return &helloWorldClient{cc}
}

func (c *helloWorldClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/helloworld.HelloWorld/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloWorldClient) GreetInManyLanguages(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (HelloWorld_GreetInManyLanguagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloWorld_ServiceDesc.Streams[0], "/helloworld.HelloWorld/GreetInManyLanguages", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldGreetInManyLanguagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloWorld_GreetInManyLanguagesClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type helloWorldGreetInManyLanguagesClient struct {
	grpc.ClientStream
}

func (x *helloWorldGreetInManyLanguagesClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloWorldClient) GreetManyPeople(ctx context.Context, opts ...grpc.CallOption) (HelloWorld_GreetManyPeopleClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloWorld_ServiceDesc.Streams[1], "/helloworld.HelloWorld/GreetManyPeople", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldGreetManyPeopleClient{stream}
	return x, nil
}

type HelloWorld_GreetManyPeopleClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type helloWorldGreetManyPeopleClient struct {
	grpc.ClientStream
}

func (x *helloWorldGreetManyPeopleClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloWorldGreetManyPeopleClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloWorldClient) Chat(ctx context.Context, opts ...grpc.CallOption) (HelloWorld_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloWorld_ServiceDesc.Streams[2], "/helloworld.HelloWorld/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloWorldChatClient{stream}
	return x, nil
}

type HelloWorld_ChatClient interface {
	Send(*Talk) error
	Recv() (*Talk, error)
	grpc.ClientStream
}

type helloWorldChatClient struct {
	grpc.ClientStream
}

func (x *helloWorldChatClient) Send(m *Talk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloWorldChatClient) Recv() (*Talk, error) {
	m := new(Talk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloWorldServer is the server API for HelloWorld service.
// All implementations must embed UnimplementedHelloWorldServer
// for forward compatibility
type HelloWorldServer interface {
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	//server-side streaming
	GreetInManyLanguages(*GreetRequest, HelloWorld_GreetInManyLanguagesServer) error
	//client-side streaming
	GreetManyPeople(HelloWorld_GreetManyPeopleServer) error
	//bidirectional streaming
	Chat(HelloWorld_ChatServer) error
	mustEmbedUnimplementedHelloWorldServer()
}

// UnimplementedHelloWorldServer must be embedded to have forward compatible implementations.
type UnimplementedHelloWorldServer struct {
}

func (UnimplementedHelloWorldServer) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedHelloWorldServer) GreetInManyLanguages(*GreetRequest, HelloWorld_GreetInManyLanguagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetInManyLanguages not implemented")
}
func (UnimplementedHelloWorldServer) GreetManyPeople(HelloWorld_GreetManyPeopleServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyPeople not implemented")
}
func (UnimplementedHelloWorldServer) Chat(HelloWorld_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedHelloWorldServer) mustEmbedUnimplementedHelloWorldServer() {}

// UnsafeHelloWorldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloWorldServer will
// result in compilation errors.
type UnsafeHelloWorldServer interface {
	mustEmbedUnimplementedHelloWorldServer()
}

func RegisterHelloWorldServer(s grpc.ServiceRegistrar, srv HelloWorldServer) {
	s.RegisterService(&HelloWorld_ServiceDesc, srv)
}

func _HelloWorld_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloWorldServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloWorld/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloWorldServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloWorld_GreetInManyLanguages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloWorldServer).GreetInManyLanguages(m, &helloWorldGreetInManyLanguagesServer{stream})
}

type HelloWorld_GreetInManyLanguagesServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type helloWorldGreetInManyLanguagesServer struct {
	grpc.ServerStream
}

func (x *helloWorldGreetInManyLanguagesServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloWorld_GreetManyPeople_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloWorldServer).GreetManyPeople(&helloWorldGreetManyPeopleServer{stream})
}

type HelloWorld_GreetManyPeopleServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type helloWorldGreetManyPeopleServer struct {
	grpc.ServerStream
}

func (x *helloWorldGreetManyPeopleServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloWorldGreetManyPeopleServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloWorld_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloWorldServer).Chat(&helloWorldChatServer{stream})
}

type HelloWorld_ChatServer interface {
	Send(*Talk) error
	Recv() (*Talk, error)
	grpc.ServerStream
}

type helloWorldChatServer struct {
	grpc.ServerStream
}

func (x *helloWorldChatServer) Send(m *Talk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloWorldChatServer) Recv() (*Talk, error) {
	m := new(Talk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloWorld_ServiceDesc is the grpc.ServiceDesc for HelloWorld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloWorld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloWorld",
	HandlerType: (*HelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _HelloWorld_Greet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetInManyLanguages",
			Handler:       _HelloWorld_GreetInManyLanguages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GreetManyPeople",
			Handler:       _HelloWorld_GreetManyPeople_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Chat",
			Handler:       _HelloWorld_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
