package handler

import (
	"context"
	"fmt"
	"github.com/0xweb-3/NewEthCEXWallet/wallet/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WalletServer struct {
	proto.UnimplementedWalletServer
}

func (s *WalletServer) BatchCreateWallet(context.Context, *proto.BatchCreateWalletReq) (*emptypb.Empty, error) {
	fmt.Println("BatchCreateWallet called")
	panic("implement me")
}

// 为新用户分配地址
func (s *WalletServer) UpdateWalletUser(context.Context, *proto.UpdateWalletUserReq) (*proto.WalletInfo, error) {
	panic("implement me")
}

// 通过用户ID查询地址地址信息
func (s *WalletServer) GetAddressByUserId(context.Context, *proto.GetAddressByUserIdReq) (*proto.WalletInfo, error) {
	panic("implement me")
}

// 通过用户地址信息查询地址地址信息
func (s *WalletServer) GetAddress(context.Context, *proto.GetAddressReq) (*proto.WalletInfo, error) {
	panic("implement me")
}
