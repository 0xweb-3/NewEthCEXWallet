package main

import (
	"context"
	"fmt"
	"github.com/0xweb-3/NewEthCEXWallet/wallet/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 1. 建立连接
	walletConn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer walletConn.Close()

	// 2. 创建 WalletSrvClient
	walletClient := proto.NewWalletClient(walletConn)

	// 3. 调用 BatchCreateWallet 方法
	_, err = walletClient.BatchCreateWallet(context.Background(), &proto.BatchCreateWalletReq{})
	if err != nil {
		panic(err)
	}
	fmt.Println()
}
