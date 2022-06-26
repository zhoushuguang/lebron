package svc

import (
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"
	"github.com/zhoushuguang/lebron/apps/seckill/rpc/internal/config"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	BizRedis    *redis.Redis
	ProductRPC  product.Product
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		BizRedis:    redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		ProductRPC:  product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		KafkaPusher: kq.NewPusher(c.Kafka.Addrs, c.Kafka.SeckillTopic),
	}
}
