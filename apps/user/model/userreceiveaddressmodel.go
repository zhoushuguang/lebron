package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	publicConstant "github.com/zhoushuguang/lebron/pkg/constant"
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
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*UserReceiveAddress, error)
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
	}

	customUserReceiveAddressModel struct {
		*defaultUserReceiveAddressModel
	}
)

func (m customUserReceiveAddressModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*UserReceiveAddress, error) {
	rowBuilder = rowBuilder.OrderBy("id DESC")

	query, values, err := rowBuilder.Where("is_delete = ?", publicConstant.IsDelNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserReceiveAddress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
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

func (m *customUserReceiveAddressModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *customUserReceiveAddressModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}
