package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserReceiveAddressModel = (*customUserReceiveAddressModel)(nil)
var (
	homestayOrderFieldNames = builder.RawFieldNames(&UserReceiveAddress{})
	homestayOrderRows       = strings.Join(homestayOrderFieldNames, ",")
)

type (
	// UserReceiveAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserReceiveAddressModel.
	UserReceiveAddressModel interface {
		userReceiveAddressModel
		UpdateIsDelete(ctx context.Context, data *UserReceiveAddress) error
		FindAllByUid(ctx context.Context, uid int64) ([]*UserReceiveAddress, error)
		RowBuilder() squirrel.SelectBuilder
	}

	customUserReceiveAddressModel struct {
		*defaultUserReceiveAddressModel
	}
)

func (m customUserReceiveAddressModel) FindAllByUid(ctx context.Context, uid int64) ([]*UserReceiveAddress, error) {
	var resp []*UserReceiveAddress
	query := fmt.Sprintf("select %s from %s where `uid` = ? and is_delete = 0", userReceiveAddressRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

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

func (m *customUserReceiveAddressModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(homestayOrderRows).From(m.table)
}
