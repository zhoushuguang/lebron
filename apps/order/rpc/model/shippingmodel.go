package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShippingModel = (*customShippingModel)(nil)

type (
	// ShippingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShippingModel.
	ShippingModel interface {
		shippingModel
	}

	customShippingModel struct {
		*defaultShippingModel
	}
)

// NewShippingModel returns a model for the database table.
func NewShippingModel(conn sqlx.SqlConn, c cache.CacheConf) ShippingModel {
	return &customShippingModel{
		defaultShippingModel: newShippingModel(conn, c),
	}
}
