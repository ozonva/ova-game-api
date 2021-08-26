// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_game_api

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HeroApiClient is the client API for HeroApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeroApiClient interface {
	CreateHero(ctx context.Context, in *CreateHeroRequest, opts ...grpc.CallOption) (*CreateGameResponse, error)
	ListHeroes(ctx context.Context, in *ListHeroRequest, opts ...grpc.CallOption) (*ListHeroResponse, error)
	DescribeHero(ctx context.Context, in *DescribeHeroRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveHero(ctx context.Context, in *RemoveHeroRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type heroApiClient struct {
	cc grpc.ClientConnInterface
}

func NewHeroApiClient(cc grpc.ClientConnInterface) HeroApiClient {
	return &heroApiClient{cc}
}

func (c *heroApiClient) CreateHero(ctx context.Context, in *CreateHeroRequest, opts ...grpc.CallOption) (*CreateGameResponse, error) {
	out := new(CreateGameResponse)
	err := c.cc.Invoke(ctx, "/ova.game.api.HeroApi/CreateHero", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroApiClient) ListHeroes(ctx context.Context, in *ListHeroRequest, opts ...grpc.CallOption) (*ListHeroResponse, error) {
	out := new(ListHeroResponse)
	err := c.cc.Invoke(ctx, "/ova.game.api.HeroApi/ListHeroes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroApiClient) DescribeHero(ctx context.Context, in *DescribeHeroRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ova.game.api.HeroApi/DescribeHero", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heroApiClient) RemoveHero(ctx context.Context, in *RemoveHeroRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ova.game.api.HeroApi/RemoveHero", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeroApiServer is the server API for HeroApi service.
// All implementations must embed UnimplementedHeroApiServer
// for forward compatibility
type HeroApiServer interface {
	CreateHero(context.Context, *CreateHeroRequest) (*CreateGameResponse, error)
	ListHeroes(context.Context, *ListHeroRequest) (*ListHeroResponse, error)
	DescribeHero(context.Context, *DescribeHeroRequest) (*empty.Empty, error)
	RemoveHero(context.Context, *RemoveHeroRequest) (*empty.Empty, error)
	mustEmbedUnimplementedHeroApiServer()
}

// UnimplementedHeroApiServer must be embedded to have forward compatible implementations.
type UnimplementedHeroApiServer struct {
}

func (UnimplementedHeroApiServer) CreateHero(context.Context, *CreateHeroRequest) (*CreateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHero not implemented")
}
func (UnimplementedHeroApiServer) ListHeroes(context.Context, *ListHeroRequest) (*ListHeroResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHeroes not implemented")
}
func (UnimplementedHeroApiServer) DescribeHero(context.Context, *DescribeHeroRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeHero not implemented")
}
func (UnimplementedHeroApiServer) RemoveHero(context.Context, *RemoveHeroRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveHero not implemented")
}
func (UnimplementedHeroApiServer) mustEmbedUnimplementedHeroApiServer() {}

// UnsafeHeroApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeroApiServer will
// result in compilation errors.
type UnsafeHeroApiServer interface {
	mustEmbedUnimplementedHeroApiServer()
}

func RegisterHeroApiServer(s grpc.ServiceRegistrar, srv HeroApiServer) {
	s.RegisterService(&HeroApi_ServiceDesc, srv)
}

func _HeroApi_CreateHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroApiServer).CreateHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.game.api.HeroApi/CreateHero",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroApiServer).CreateHero(ctx, req.(*CreateHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroApi_ListHeroes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroApiServer).ListHeroes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.game.api.HeroApi/ListHeroes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroApiServer).ListHeroes(ctx, req.(*ListHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroApi_DescribeHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroApiServer).DescribeHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.game.api.HeroApi/DescribeHero",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroApiServer).DescribeHero(ctx, req.(*DescribeHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeroApi_RemoveHero_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveHeroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeroApiServer).RemoveHero(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.game.api.HeroApi/RemoveHero",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeroApiServer).RemoveHero(ctx, req.(*RemoveHeroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeroApi_ServiceDesc is the grpc.ServiceDesc for HeroApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeroApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.game.api.HeroApi",
	HandlerType: (*HeroApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHero",
			Handler:    _HeroApi_CreateHero_Handler,
		},
		{
			MethodName: "ListHeroes",
			Handler:    _HeroApi_ListHeroes_Handler,
		},
		{
			MethodName: "DescribeHero",
			Handler:    _HeroApi_DescribeHero_Handler,
		},
		{
			MethodName: "RemoveHero",
			Handler:    _HeroApi_RemoveHero_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hero.proto",
}
