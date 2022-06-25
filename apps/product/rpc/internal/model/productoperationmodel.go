package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductOperationModel = (*customProductOperationModel)(nil)

type (
	// ProductOperationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductOperationModel.
	ProductOperationModel interface {
		productOperationModel
		OperationProducts(ctx context.Context, status int64) ([]*ProductOperation, error)
	}

	customProductOperationModel struct {
		*defaultProductOperationModel
	}
)

// NewProductOperationModel returns a model for the database table.
func NewProductOperationModel(conn sqlx.SqlConn, c cache.CacheConf) ProductOperationModel {
	return &customProductOperationModel{
		defaultProductOperationModel: newProductOperationModel(conn, c),
	}
}

func (m *customProductOperationModel) OperationProducts(ctx context.Context, status int64) ([]*ProductOperation, error) {
	var operations []*ProductOperation
	err := m.QueryRowsNoCacheCtx(ctx, &operations, fmt.Sprintf("select %s from %s where status=?", productOperationRows, m.table), status)
	if err != nil {
		return nil, err
	}
	return operations, err
}
