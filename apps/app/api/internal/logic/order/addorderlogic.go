package order

import (
	"context"
	"encoding/json"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
	"github.com/zhoushuguang/lebron/apps/order/rpc/order"
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOrderLogic) AddOrder(req *types.OrderAddReq) (resp *types.OrderAddResp, err error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	// 获取 OrderRpc BuildTarget
	orderRpcServer, err := l.svcCtx.Config.OrderRPC.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// 获取 ProductRpc BuildTarget
	productRpcBusiServer, err := l.svcCtx.Config.ProductRPC.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// dtm 服务的 etcd 注册地址
	//var dtmServer = "etcd://127.0.0.1:2379/dtmservice"
	var dtmServer = "discov://127.0.0.1:2379/dtmservice"
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcServer+"/order.Order/CreateOrderDTM",
			orderRpcServer+"/order.Order/CreateOrderDTMRevert",
			&order.AddOrderReq{
				Userid:           uid,
				Productid:        req.Productid,
				Quantity:         req.Quantity,
				Postage:          req.Postage,
				ReceiveAddressId: req.ReceiveAddressId,
			}).
		Add(productRpcBusiServer+"/product.Product/DecrStock",
			productRpcBusiServer+"/product.Product/DecrStockRevert",
			&product.DecrStockRequest{
				Id:  req.Productid,
				Num: req.Quantity,
			})

	// 事务提交
	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.OrderAddResp{
		Id: "test",
	}, nil
}
