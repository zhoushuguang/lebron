package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		CategoryProducts(ctx context.Context, ctime string, cateid, limit int64) ([]*Product, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c),
	}
}

func (m *customProductModel) CategoryProducts(ctx context.Context, ctime string, cateid, limit int64) ([]*Product, error) {
	var products []*Product
	err := m.QueryRowsNoCacheCtx(ctx, &products, fmt.Sprintf("select %s from %s where cateid=? and status=1 and create_time<? order by create_time desc limit ?", productRows, m.table), cateid, ctime, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}
