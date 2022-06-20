// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user.proto

package user

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 登录
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// 获取用户信息
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	// 添加收获地址
	AddUserReceiveAddress(ctx context.Context, in *UserReceiveAddressAddReq, opts ...grpc.CallOption) (*UserReceiveAddressAddRes, error)
	// 编辑收获地址
	EditUserReceiveAddress(ctx context.Context, in *UserReceiveAddressEditReq, opts ...grpc.CallOption) (*UserReceiveAddressEditRes, error)
	// 删除收获地址
	DelUserReceiveAddress(ctx context.Context, in *UserReceiveAddressDelReq, opts ...grpc.CallOption) (*UserReceiveAddressDelRes, error)
	// 获取收获地址列表
	GetUserReceiveAddressList(ctx context.Context, in *UserReceiveAddressListReq, opts ...grpc.CallOption) (*UserReceiveAddressListRes, error)
	// 添加收藏
	AddCollection(ctx context.Context, in *CollectionAddReq, opts ...grpc.CallOption) (*CollectionAddRes, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/user.User/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, "/user.User/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddUserReceiveAddress(ctx context.Context, in *UserReceiveAddressAddReq, opts ...grpc.CallOption) (*UserReceiveAddressAddRes, error) {
	out := new(UserReceiveAddressAddRes)
	err := c.cc.Invoke(ctx, "/user.User/addUserReceiveAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) EditUserReceiveAddress(ctx context.Context, in *UserReceiveAddressEditReq, opts ...grpc.CallOption) (*UserReceiveAddressEditRes, error) {
	out := new(UserReceiveAddressEditRes)
	err := c.cc.Invoke(ctx, "/user.User/editUserReceiveAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DelUserReceiveAddress(ctx context.Context, in *UserReceiveAddressDelReq, opts ...grpc.CallOption) (*UserReceiveAddressDelRes, error) {
	out := new(UserReceiveAddressDelRes)
	err := c.cc.Invoke(ctx, "/user.User/delUserReceiveAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserReceiveAddressList(ctx context.Context, in *UserReceiveAddressListReq, opts ...grpc.CallOption) (*UserReceiveAddressListRes, error) {
	out := new(UserReceiveAddressListRes)
	err := c.cc.Invoke(ctx, "/user.User/getUserReceiveAddressList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddCollection(ctx context.Context, in *CollectionAddReq, opts ...grpc.CallOption) (*CollectionAddRes, error) {
	out := new(CollectionAddRes)
	err := c.cc.Invoke(ctx, "/user.User/addCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// 登录
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// 获取用户信息
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	// 添加收获地址
	AddUserReceiveAddress(context.Context, *UserReceiveAddressAddReq) (*UserReceiveAddressAddRes, error)
	// 编辑收获地址
	EditUserReceiveAddress(context.Context, *UserReceiveAddressEditReq) (*UserReceiveAddressEditRes, error)
	// 删除收获地址
	DelUserReceiveAddress(context.Context, *UserReceiveAddressDelReq) (*UserReceiveAddressDelRes, error)
	// 获取收获地址列表
	GetUserReceiveAddressList(context.Context, *UserReceiveAddressListReq) (*UserReceiveAddressListRes, error)
	// 添加收藏
	AddCollection(context.Context, *CollectionAddReq) (*CollectionAddRes, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServer) AddUserReceiveAddress(context.Context, *UserReceiveAddressAddReq) (*UserReceiveAddressAddRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserReceiveAddress not implemented")
}
func (UnimplementedUserServer) EditUserReceiveAddress(context.Context, *UserReceiveAddressEditReq) (*UserReceiveAddressEditRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserReceiveAddress not implemented")
}
func (UnimplementedUserServer) DelUserReceiveAddress(context.Context, *UserReceiveAddressDelReq) (*UserReceiveAddressDelRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUserReceiveAddress not implemented")
}
func (UnimplementedUserServer) GetUserReceiveAddressList(context.Context, *UserReceiveAddressListReq) (*UserReceiveAddressListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserReceiveAddressList not implemented")
}
func (UnimplementedUserServer) AddCollection(context.Context, *CollectionAddReq) (*CollectionAddRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollection not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddUserReceiveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReceiveAddressAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddUserReceiveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/addUserReceiveAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddUserReceiveAddress(ctx, req.(*UserReceiveAddressAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_EditUserReceiveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReceiveAddressEditReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).EditUserReceiveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/editUserReceiveAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).EditUserReceiveAddress(ctx, req.(*UserReceiveAddressEditReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DelUserReceiveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReceiveAddressDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DelUserReceiveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/delUserReceiveAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DelUserReceiveAddress(ctx, req.(*UserReceiveAddressDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserReceiveAddressList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReceiveAddressListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserReceiveAddressList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUserReceiveAddressList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserReceiveAddressList(ctx, req.(*UserReceiveAddressListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/addCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddCollection(ctx, req.(*CollectionAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _User_UserInfo_Handler,
		},
		{
			MethodName: "addUserReceiveAddress",
			Handler:    _User_AddUserReceiveAddress_Handler,
		},
		{
			MethodName: "editUserReceiveAddress",
			Handler:    _User_EditUserReceiveAddress_Handler,
		},
		{
			MethodName: "delUserReceiveAddress",
			Handler:    _User_DelUserReceiveAddress_Handler,
		},
		{
			MethodName: "getUserReceiveAddressList",
			Handler:    _User_GetUserReceiveAddressList_Handler,
		},
		{
			MethodName: "addCollection",
			Handler:    _User_AddCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
