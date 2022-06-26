package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	BizRedis   redis.RedisConf
	ProductRPC zrpc.RpcClientConf
	Kafka      struct {
		Addrs        []string
		SeckillTopic string
	}
}
