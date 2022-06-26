package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShoppingModel = (*customShoppingModel)(nil)

type (
	// ShoppingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShoppingModel.
	ShoppingModel interface {
		shoppingModel
	}

	customShoppingModel struct {
		*defaultShoppingModel
	}
)

// NewShoppingModel returns a model for the database table.
func NewShoppingModel(conn sqlx.SqlConn, c cache.CacheConf) ShoppingModel {
	return &customShoppingModel{
		defaultShoppingModel: newShoppingModel(conn, c),
	}
}
