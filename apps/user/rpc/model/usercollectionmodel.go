package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserCollectionModel = (*customUserCollectionModel)(nil)

type (
	// UserCollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCollectionModel.
	UserCollectionModel interface {
		userCollectionModel
		FindAllByUid(ctx context.Context, uid int64) ([]*UserCollection, error)
		UpdateIsDelete(ctx context.Context, data *UserCollection) error
	}

	customUserCollectionModel struct {
		*defaultUserCollectionModel
	}
)

func (m customUserCollectionModel) UpdateIsDelete(ctx context.Context, data *UserCollection) error {
	userCollectionIdKey := fmt.Sprintf("%s%v", cacheUserCollectionIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set is_delete = 1 where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, data.Id)
	}, userCollectionIdKey)
	return err
}

func (m customUserCollectionModel) FindAllByUid(ctx context.Context, uid int64) ([]*UserCollection, error) {
	var resp []*UserCollection
	query := fmt.Sprintf("select %s from %s where `uid` = ? and is_delete = 0", userCollectionRows, m.table)
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

// NewUserCollectionModel returns a model for the database table.
func NewUserCollectionModel(conn sqlx.SqlConn, c cache.CacheConf) UserCollectionModel {
	return &customUserCollectionModel{
		defaultUserCollectionModel: newUserCollectionModel(conn, c),
	}
}
