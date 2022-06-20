package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhoushuguang/lebron/apps/user/model"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//add a dependency on user model
	UserModel       model.UserModel
	CollectionModel model.CollectionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserModel:       model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		CollectionModel: model.NewCollectionModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
