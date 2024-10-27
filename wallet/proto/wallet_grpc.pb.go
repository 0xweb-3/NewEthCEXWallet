// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.7
// source: wallet.proto

package proto

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

const (
	Wallet_BatchCreateWallet_FullMethodName  = "/Wallet/BatchCreateWallet"
	Wallet_UpdateWalletUser_FullMethodName   = "/Wallet/UpdateWalletUser"
	Wallet_GetAddressByUserId_FullMethodName = "/Wallet/GetAddressByUserId"
	Wallet_GetAddress_FullMethodName         = "/Wallet/GetAddress"
)

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletClient interface {
	// 批量生成地址
	BatchCreateWallet(ctx context.Context, in *BatchCreateWalletReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 通过私钥生成地址信息
	//
	//	rpc CreateAddressByPrivateKey(CreateAddressByPrivateKeyReq) returns (google.protobuf.Empty);
	//
	// 将公钥转换为地址信息
	//
	//	rpc GetAddressByPublicKey(GetAddressByPublicKeyReq) returns(google.protobuf.Empty);
	//
	// 为新用户分配地址
	UpdateWalletUser(ctx context.Context, in *UpdateWalletUserReq, opts ...grpc.CallOption) (*WalletInfo, error)
	// 通过用户ID查询地址地址信息
	GetAddressByUserId(ctx context.Context, in *GetAddressByUserIdReq, opts ...grpc.CallOption) (*WalletInfo, error)
	// 通过用户地址信息查询地址地址信息
	GetAddress(ctx context.Context, in *GetAddressReq, opts ...grpc.CallOption) (*WalletInfo, error)
}

type walletClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletClient(cc grpc.ClientConnInterface) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) BatchCreateWallet(ctx context.Context, in *BatchCreateWalletReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Wallet_BatchCreateWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) UpdateWalletUser(ctx context.Context, in *UpdateWalletUserReq, opts ...grpc.CallOption) (*WalletInfo, error) {
	out := new(WalletInfo)
	err := c.cc.Invoke(ctx, Wallet_UpdateWalletUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAddressByUserId(ctx context.Context, in *GetAddressByUserIdReq, opts ...grpc.CallOption) (*WalletInfo, error) {
	out := new(WalletInfo)
	err := c.cc.Invoke(ctx, Wallet_GetAddressByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetAddress(ctx context.Context, in *GetAddressReq, opts ...grpc.CallOption) (*WalletInfo, error) {
	out := new(WalletInfo)
	err := c.cc.Invoke(ctx, Wallet_GetAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
// All implementations must embed UnimplementedWalletServer
// for forward compatibility
type WalletServer interface {
	// 批量生成地址
	BatchCreateWallet(context.Context, *BatchCreateWalletReq) (*emptypb.Empty, error)
	// 通过私钥生成地址信息
	//
	//	rpc CreateAddressByPrivateKey(CreateAddressByPrivateKeyReq) returns (google.protobuf.Empty);
	//
	// 将公钥转换为地址信息
	//
	//	rpc GetAddressByPublicKey(GetAddressByPublicKeyReq) returns(google.protobuf.Empty);
	//
	// 为新用户分配地址
	UpdateWalletUser(context.Context, *UpdateWalletUserReq) (*WalletInfo, error)
	// 通过用户ID查询地址地址信息
	GetAddressByUserId(context.Context, *GetAddressByUserIdReq) (*WalletInfo, error)
	// 通过用户地址信息查询地址地址信息
	GetAddress(context.Context, *GetAddressReq) (*WalletInfo, error)
	mustEmbedUnimplementedWalletServer()
}

// UnimplementedWalletServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (UnimplementedWalletServer) BatchCreateWallet(context.Context, *BatchCreateWalletReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateWallet not implemented")
}
func (UnimplementedWalletServer) UpdateWalletUser(context.Context, *UpdateWalletUserReq) (*WalletInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWalletUser not implemented")
}
func (UnimplementedWalletServer) GetAddressByUserId(context.Context, *GetAddressByUserIdReq) (*WalletInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressByUserId not implemented")
}
func (UnimplementedWalletServer) GetAddress(context.Context, *GetAddressReq) (*WalletInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedWalletServer) mustEmbedUnimplementedWalletServer() {}

// UnsafeWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServer will
// result in compilation errors.
type UnsafeWalletServer interface {
	mustEmbedUnimplementedWalletServer()
}

func RegisterWalletServer(s grpc.ServiceRegistrar, srv WalletServer) {
	s.RegisterService(&Wallet_ServiceDesc, srv)
}

func _Wallet_BatchCreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateWalletReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).BatchCreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_BatchCreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).BatchCreateWallet(ctx, req.(*BatchCreateWalletReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_UpdateWalletUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWalletUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).UpdateWalletUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_UpdateWalletUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).UpdateWalletUser(ctx, req.(*UpdateWalletUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAddressByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAddressByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetAddressByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAddressByUserId(ctx, req.(*GetAddressByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Wallet_GetAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetAddress(ctx, req.(*GetAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Wallet_ServiceDesc is the grpc.ServiceDesc for Wallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchCreateWallet",
			Handler:    _Wallet_BatchCreateWallet_Handler,
		},
		{
			MethodName: "UpdateWalletUser",
			Handler:    _Wallet_UpdateWalletUser_Handler,
		},
		{
			MethodName: "GetAddressByUserId",
			Handler:    _Wallet_GetAddressByUserId_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _Wallet_GetAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}