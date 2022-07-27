package logic

import (
	"context"
	"database/sql"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhoushuguang/lebron/apps/product/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockRevertLogic {
	return &DecrStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 扣减库存,再加回去
func (l *DecrStockRevertLogic) DecrStockRevert(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
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
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 更新产品库存
		_, err := l.svcCtx.ProductModel.TxUpdateStock(tx, in.Id, 1)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &product.DecrStockResponse{}, nil
}
