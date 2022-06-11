package svc

import (
	"github.com/zhoushuguang/lebron/apps/app/api/internal/config"
	"github.com/zhoushuguang/lebron/apps/order/rpc/order"
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"
	"github.com/zhoushuguang/lebron/apps/reply/rpc/reply"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	OrderRPC order.Order
	ProductRPC product.Product
	ReplyRPC reply.Reply
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OrderRPC: order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		ReplyRPC: reply.NewReply(zrpc.MustNewClient(c.ReplyRPC)),
	}
}
