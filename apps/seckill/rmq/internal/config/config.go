package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Kafka      kq.KqConf
	ProductRPC zrpc.RpcClientConf
	OrderRPC   zrpc.RpcClientConf
}
