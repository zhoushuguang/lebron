package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhoushuguang/lebron/apps/order/rpc/internal/config"
	"github.com/zhoushuguang/lebron/apps/order/rpc/internal/server"
	"github.com/zhoushuguang/lebron/apps/order/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/order/rpc/order"
	"github.com/zhoushuguang/lebron/pkg/snowflake"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/order.yaml", "the etc file")

func main() {

	orderId := snowflake.GenIDString()
	logx.Info(orderId)
	flag.Parse()

	//close statis log
	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewOrderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		order.RegisterOrderServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
