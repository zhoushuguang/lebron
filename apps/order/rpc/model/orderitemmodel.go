package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderitemModel = (*customOrderitemModel)(nil)

type (
	// OrderitemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderitemModel.
	OrderitemModel interface {
		orderitemModel
	}

	customOrderitemModel struct {
		*defaultOrderitemModel
	}
)

// NewOrderitemModel returns a model for the database table.
func NewOrderitemModel(conn sqlx.SqlConn, c cache.CacheConf) OrderitemModel {
	return &customOrderitemModel{
		defaultOrderitemModel: newOrderitemModel(conn, c),
	}
}
