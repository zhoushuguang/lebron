package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CollectionModel = (*customCollectionModel)(nil)

type (
	// CollectionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollectionModel.
	CollectionModel interface {
		collectionModel
	}

	customCollectionModel struct {
		*defaultCollectionModel
	}
)

// NewCollectionModel returns a model for the database table.
func NewCollectionModel(conn sqlx.SqlConn, c cache.CacheConf) CollectionModel {
	return &customCollectionModel{
		defaultCollectionModel: newCollectionModel(conn, c),
	}
}
