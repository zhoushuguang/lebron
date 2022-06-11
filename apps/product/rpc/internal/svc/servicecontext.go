package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhoushuguang/lebron/apps/product/rpc/internal/config"
	"github.com/zhoushuguang/lebron/apps/product/rpc/internal/model"
)

type ServiceContext struct {
	Config config.Config
	ProductModel model.ProductModel
	CategoryModel model.CategoryModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config: c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
		CategoryModel: model.NewCategoryModel(conn, c.CacheRedis),
	}
}
