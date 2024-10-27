package main

import (
	"github.com/0xweb-3/NewEthCEXWallet/wallet/handler"
	"github.com/0xweb-3/NewEthCEXWallet/wallet/proto"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. 实例化一个 server
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	// 2. 注册处理逻辑 handler
	proto.RegisterWalletServer(server, &handler.WalletServer{}) // 注册 Wallet 服务

	// 3. 启动服务
	go func() {
		err = server.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	// 接受服务退出信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
