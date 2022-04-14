// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/cart.proto

package cart

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

// CartClient is the client API for Cart service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartClient interface {
	GetSingleCart(ctx context.Context, in *CartInfo, opts ...grpc.CallOption) (*CartData, error)
	ClearCart(ctx context.Context, in *CartInfo, opts ...grpc.CallOption) (*CartData, error)
}

type cartClient struct {
	cc grpc.ClientConnInterface
}

func NewCartClient(cc grpc.ClientConnInterface) CartClient {
	return &cartClient{cc}
}

func (c *cartClient) GetSingleCart(ctx context.Context, in *CartInfo, opts ...grpc.CallOption) (*CartData, error) {
	out := new(CartData)
	err := c.cc.Invoke(ctx, "/cart.Cart/GetSingleCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartClient) ClearCart(ctx context.Context, in *CartInfo, opts ...grpc.CallOption) (*CartData, error) {
	out := new(CartData)
	err := c.cc.Invoke(ctx, "/cart.Cart/ClearCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServer is the server API for Cart service.
// All implementations must embed UnimplementedCartServer
// for forward compatibility
type CartServer interface {
	GetSingleCart(context.Context, *CartInfo) (*CartData, error)
	ClearCart(context.Context, *CartInfo) (*CartData, error)
	mustEmbedUnimplementedCartServer()
}

// UnimplementedCartServer must be embedded to have forward compatible implementations.
type UnimplementedCartServer struct {
}

func (UnimplementedCartServer) GetSingleCart(context.Context, *CartInfo) (*CartData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleCart not implemented")
}
func (UnimplementedCartServer) ClearCart(context.Context, *CartInfo) (*CartData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedCartServer) mustEmbedUnimplementedCartServer() {}

// UnsafeCartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServer will
// result in compilation errors.
type UnsafeCartServer interface {
	mustEmbedUnimplementedCartServer()
}

func RegisterCartServer(s grpc.ServiceRegistrar, srv CartServer) {
	s.RegisterService(&Cart_ServiceDesc, srv)
}

func _Cart_GetSingleCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).GetSingleCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.Cart/GetSingleCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).GetSingleCart(ctx, req.(*CartInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cart_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.Cart/ClearCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServer).ClearCart(ctx, req.(*CartInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Cart_ServiceDesc is the grpc.ServiceDesc for Cart service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cart_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.Cart",
	HandlerType: (*CartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSingleCart",
			Handler:    _Cart_GetSingleCart_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _Cart_ClearCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cart.proto",
}
