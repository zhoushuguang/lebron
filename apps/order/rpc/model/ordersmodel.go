package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrdersModel = (*customOrdersModel)(nil)

type (
	// OrdersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrdersModel.
	OrdersModel interface {
		ordersModel
		FindOneByUid(ctx context.Context, uid int64) (*Orders, error)
		CreateOrder(ctx context.Context, oid string, uid, pid int64) error
		UpdateOrderStatus(ctx context.Context, oid string, status int) error
		TxInsert(tx *sql.Tx, data *Orders) (sql.Result, error)
		TxUpdate(tx *sql.Tx, data *Orders) error
	}

	customOrdersModel struct {
		*defaultOrdersModel
	}
)

// NewOrdersModel returns a model for the database table.
func NewOrdersModel(conn sqlx.SqlConn, c cache.CacheConf) OrdersModel {
	return &customOrdersModel{
		defaultOrdersModel: newOrdersModel(conn, c),
	}
}

func (m *customOrdersModel) CreateOrder(ctx context.Context, oid string, uid, pid int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		err := conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
			_, err := session.ExecCtx(ctx, "INSERT INTO orders(id, userid) VALUES(?,?)", oid, uid)
			if err != nil {
				return err
			}
			_, err = session.ExecCtx(ctx, "INSERT INTO orderitem(orderid, userid, proid) VALUES(?,?,?)", "", uid, pid)
			return err
		})
		return nil, err
	})
	return err
}

func (m *customOrdersModel) UpdateOrderStatus(ctx context.Context, oid string, status int) error {
	ordersOrdersIdKey := fmt.Sprintf("%s%v", cacheOrdersIdPrefix, oid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET status = ? WHERE id = ?", m.table), status, oid)
	}, ordersOrdersIdKey)
	return err
}

//func (m *defaultOrdersModel) TxInsert(tx *sql.Tx, data *Orders) (sql.Result, error) {
func (m *customOrdersModel) TxInsert(tx *sql.Tx, data *Orders) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, ordersRowsExpectAutoSet)
	ret, err := tx.Exec(query, data.Id, data.Userid, data.Shoppingid, data.Payment, data.Paymenttype, data.Postage, data.Status)
	return ret, err
}

func (m *customOrdersModel) FindOneByUid(ctx context.Context, uid int64) (*Orders, error) {
	var resp Orders

	query := fmt.Sprintf("select %s from %s where `uid` = ? order by create_time desc limit 1", ordersRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customOrdersModel) TxUpdate(tx *sql.Tx, data *Orders) error {
	productIdKey := fmt.Sprintf("%s%v", cacheOrdersIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ordersRowsWithPlaceHolder)
		return tx.Exec(query, data.Userid, data.Shoppingid, data.Payment, data.Paymenttype, data.Postage, data.Status, data.Id)
	}, productIdKey)
	return err
}
