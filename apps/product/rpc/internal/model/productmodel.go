package model

import (
	"context"
	"database/sql"
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
		UpdateProductStock(ctx context.Context, pid, num int64) error
		TxUpdateStock(tx *sql.Tx, id int64, delta int) (sql.Result, error)
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

func (m *customProductModel) UpdateProductStock(ctx context.Context, pid, num int64) error {
	productProductIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, pid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET stock = stock - ? WHERE id = ? and stock > 0", m.table), num, pid)
	}, productProductIdKey)
	return err
}

func (m *customProductModel) TxUpdateStock(tx *sql.Tx, id int64, delta int) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheProductProductIdPrefix, id)
	return m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set stock = stock + ? where stock >= -? and id=?", m.table)
		return tx.Exec(query, delta, delta, id)
	}, productIdKey)
}
