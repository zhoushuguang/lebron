package order

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
	"github.com/zhoushuguang/lebron/apps/order/rpc/order"

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
	addOrderReq := order.AddOrderReq{
		Userid:           uid,
		Productid:        req.Productid,
		Quantity:         req.Quantity,
		Postage:          req.Postage,
		ReceiveAddressId: req.ReceiveAddressId,
	}
	addOrderResp, err := l.svcCtx.OrderRPC.CreateOrderDTM(l.ctx, &addOrderReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.OrderAddResp{
		Id: addOrderResp.Id,
	}, nil
}
