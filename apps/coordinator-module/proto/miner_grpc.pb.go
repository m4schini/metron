// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// MinerClient is the client API for Miner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MinerClient interface {
	GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Account, error)
	GetVideoDetails(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (*VideoDetails, error)
}

type minerClient struct {
	cc grpc.ClientConnInterface
}

func NewMinerClient(cc grpc.ClientConnInterface) MinerClient {
	return &minerClient{cc}
}

func (c *minerClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/miner.Miner/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) GetVideoDetails(ctx context.Context, in *VideoRequest, opts ...grpc.CallOption) (*VideoDetails, error) {
	out := new(VideoDetails)
	err := c.cc.Invoke(ctx, "/miner.Miner/GetVideoDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MinerServer is the server API for Miner service.
// All implementations must embed UnimplementedMinerServer
// for forward compatibility
type MinerServer interface {
	GetAccount(context.Context, *AccountRequest) (*Account, error)
	GetVideoDetails(context.Context, *VideoRequest) (*VideoDetails, error)
	mustEmbedUnimplementedMinerServer()
}

// UnimplementedMinerServer must be embedded to have forward compatible implementations.
type UnimplementedMinerServer struct {
}

func (UnimplementedMinerServer) GetAccount(context.Context, *AccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedMinerServer) GetVideoDetails(context.Context, *VideoRequest) (*VideoDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoDetails not implemented")
}
func (UnimplementedMinerServer) mustEmbedUnimplementedMinerServer() {}

// UnsafeMinerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MinerServer will
// result in compilation errors.
type UnsafeMinerServer interface {
	mustEmbedUnimplementedMinerServer()
}

func RegisterMinerServer(s grpc.ServiceRegistrar, srv MinerServer) {
	s.RegisterService(&Miner_ServiceDesc, srv)
}

func _Miner_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miner.Miner/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).GetAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_GetVideoDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).GetVideoDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miner.Miner/GetVideoDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).GetVideoDetails(ctx, req.(*VideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Miner_ServiceDesc is the grpc.ServiceDesc for Miner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Miner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "miner.Miner",
	HandlerType: (*MinerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _Miner_GetAccount_Handler,
		},
		{
			MethodName: "GetVideoDetails",
			Handler:    _Miner_GetVideoDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "miner.proto",
}