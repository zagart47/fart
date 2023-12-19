// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/fart.proto

package proto

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

// FartClient is the client API for Fart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FartClient interface {
	DoCmd(ctx context.Context, in *CmdRequest, opts ...grpc.CallOption) (*CmdResponse, error)
}

type fartClient struct {
	cc grpc.ClientConnInterface
}

func NewFartClient(cc grpc.ClientConnInterface) FartClient {
	return &fartClient{cc}
}

func (c *fartClient) DoCmd(ctx context.Context, in *CmdRequest, opts ...grpc.CallOption) (*CmdResponse, error) {
	out := new(CmdResponse)
	err := c.cc.Invoke(ctx, "/proto.Fart/DoCmd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FartServer is the server API for Fart service.
// All implementations must embed UnimplementedFartServer
// for forward compatibility
type FartServer interface {
	DoCmd(context.Context, *CmdRequest) (*CmdResponse, error)
	mustEmbedUnimplementedFartServer()
}

// UnimplementedFartServer must be embedded to have forward compatible implementations.
type UnimplementedFartServer struct {
}

func (UnimplementedFartServer) DoCmd(context.Context, *CmdRequest) (*CmdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoCmd not implemented")
}
func (UnimplementedFartServer) mustEmbedUnimplementedFartServer() {}

// UnsafeFartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FartServer will
// result in compilation errors.
type UnsafeFartServer interface {
	mustEmbedUnimplementedFartServer()
}

func RegisterFartServer(s grpc.ServiceRegistrar, srv FartServer) {
	s.RegisterService(&Fart_ServiceDesc, srv)
}

func _Fart_DoCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FartServer).DoCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Fart/DoCmd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FartServer).DoCmd(ctx, req.(*CmdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fart_ServiceDesc is the grpc.ServiceDesc for Fart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Fart",
	HandlerType: (*FartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoCmd",
			Handler:    _Fart_DoCmd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/fart.proto",
}
