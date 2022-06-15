package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserReceiveAddressModel = (*customUserReceiveAddressModel)(nil)

type (
	// UserReceiveAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserReceiveAddressModel.
	UserReceiveAddressModel interface {
		userReceiveAddressModel
	}

	customUserReceiveAddressModel struct {
		*defaultUserReceiveAddressModel
	}
)

// NewUserReceiveAddressModel returns a model for the database table.
func NewUserReceiveAddressModel(conn sqlx.SqlConn, c cache.CacheConf) UserReceiveAddressModel {
	return &customUserReceiveAddressModel{
		defaultUserReceiveAddressModel: newUserReceiveAddressModel(conn, c),
	}
}
