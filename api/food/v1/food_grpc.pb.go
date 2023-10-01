// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: food/v1/food.proto

package v1

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
	Food_LoginByUsername_FullMethodName    = "/api.food.v1.Food/LoginByUsername"
	Food_ListCollectionShop_FullMethodName = "/api.food.v1.Food/ListCollectionShop"
)

// FoodClient is the client API for Food service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FoodClient interface {
	LoginByUsername(ctx context.Context, in *LoginByUsernameRequest, opts ...grpc.CallOption) (*LoginByUsernameReply, error)
	ListCollectionShop(ctx context.Context, in *ListCollectionShopRequest, opts ...grpc.CallOption) (*ListCollectionShopReply, error)
}

type foodClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodClient(cc grpc.ClientConnInterface) FoodClient {
	return &foodClient{cc}
}

func (c *foodClient) LoginByUsername(ctx context.Context, in *LoginByUsernameRequest, opts ...grpc.CallOption) (*LoginByUsernameReply, error) {
	out := new(LoginByUsernameReply)
	err := c.cc.Invoke(ctx, Food_LoginByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) ListCollectionShop(ctx context.Context, in *ListCollectionShopRequest, opts ...grpc.CallOption) (*ListCollectionShopReply, error) {
	out := new(ListCollectionShopReply)
	err := c.cc.Invoke(ctx, Food_ListCollectionShop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoodServer is the server API for Food service.
// All implementations must embed UnimplementedFoodServer
// for forward compatibility
type FoodServer interface {
	LoginByUsername(context.Context, *LoginByUsernameRequest) (*LoginByUsernameReply, error)
	ListCollectionShop(context.Context, *ListCollectionShopRequest) (*ListCollectionShopReply, error)
	mustEmbedUnimplementedFoodServer()
}

// UnimplementedFoodServer must be embedded to have forward compatible implementations.
type UnimplementedFoodServer struct {
}

func (UnimplementedFoodServer) LoginByUsername(context.Context, *LoginByUsernameRequest) (*LoginByUsernameReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByUsername not implemented")
}
func (UnimplementedFoodServer) ListCollectionShop(context.Context, *ListCollectionShopRequest) (*ListCollectionShopReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollectionShop not implemented")
}
func (UnimplementedFoodServer) mustEmbedUnimplementedFoodServer() {}

// UnsafeFoodServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FoodServer will
// result in compilation errors.
type UnsafeFoodServer interface {
	mustEmbedUnimplementedFoodServer()
}

func RegisterFoodServer(s grpc.ServiceRegistrar, srv FoodServer) {
	s.RegisterService(&Food_ServiceDesc, srv)
}

func _Food_LoginByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).LoginByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Food_LoginByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).LoginByUsername(ctx, req.(*LoginByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_ListCollectionShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).ListCollectionShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Food_ListCollectionShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).ListCollectionShop(ctx, req.(*ListCollectionShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Food_ServiceDesc is the grpc.ServiceDesc for Food service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Food_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.food.v1.Food",
	HandlerType: (*FoodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginByUsername",
			Handler:    _Food_LoginByUsername_Handler,
		},
		{
			MethodName: "ListCollectionShop",
			Handler:    _Food_ListCollectionShop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "food/v1/food.proto",
}
