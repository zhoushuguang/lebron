package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	OrderRPC zrpc.RpcClientConf
	ProductRPC zrpc.RpcClientConf
	ReplyRPC zrpc.RpcClientConf
}
