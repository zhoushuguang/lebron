package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserReceiveAddressModel = (*customUserReceiveAddressModel)(nil)

type (
	// UserReceiveAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserReceiveAddressModel.
	UserReceiveAddressModel interface {
		userReceiveAddressModel
		UpdateIsDelete(ctx context.Context, data *UserReceiveAddress) error
	}

	customUserReceiveAddressModel struct {
		*defaultUserReceiveAddressModel
	}
)

func (m customUserReceiveAddressModel) UpdateIsDelete(ctx context.Context, data *UserReceiveAddress) error {
	userReceiveAddressIdKey := fmt.Sprintf("%s%v", cacheUserReceiveAddressIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set is_delete = 1 where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, data.Id)
	}, userReceiveAddressIdKey)
	return err
}

// NewUserReceiveAddressModel returns a model for the database table.
func NewUserReceiveAddressModel(conn sqlx.SqlConn, c cache.CacheConf) UserReceiveAddressModel {
	return &customUserReceiveAddressModel{
		defaultUserReceiveAddressModel: newUserReceiveAddressModel(conn, c),
	}
}
