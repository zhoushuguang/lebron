package logic

import (
	"context"
	"database/sql"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhoushuguang/lebron/apps/product/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockLogic {
	return &DecrStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockLogic) DecrStock(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
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
		result, err := l.svcCtx.ProductModel.TxUpdateStock(tx, in.Id, -1)
		if err != nil {
			return err
		}

		affected, err := result.RowsAffected()
		// 库存不足，返回子事务失败
		if err == nil && affected == 0 {
			return dtmcli.ErrFailure
		}

		return err
	})

	// 库存不足，不再重试，走回滚
	if err == dtmcli.ErrFailure {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	if err != nil {
		return nil, err
	}

	return &product.DecrStockResponse{}, nil
}
