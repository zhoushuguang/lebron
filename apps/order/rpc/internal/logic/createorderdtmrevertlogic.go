package logic

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhoushuguang/lebron/apps/order/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/order/rpc/order"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderDTMRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderDTMRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderDTMRevertLogic {
	return &CreateOrderDTMRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderDTMRevertLogic) CreateOrderDTMRevert(in *order.AddOrderReq) (*order.AddOrderResp, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 查询用户是否存在
		_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
			Id: in.Userid,
		})
		if err != nil {
			return fmt.Errorf("用户不存在")
		}
		// 查询用户最新创建的订单
		resOrder, err := l.svcCtx.OrderModel.FindOneByUid(l.ctx, in.Userid)
		if err != nil {
			return fmt.Errorf("订单不存在")
		}
		// 修改订单状态60，标识订单 已关闭
		resOrder.Status = 60
		err = l.svcCtx.OrderModel.TxUpdate(tx, resOrder)
		if err != nil {
			return fmt.Errorf("订单更新失败")
		}

		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &order.AddOrderResp{}, nil
}
