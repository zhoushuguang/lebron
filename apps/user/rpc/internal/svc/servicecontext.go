package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/config"
	"github.com/zhoushuguang/lebron/apps/user/rpc/model"
)

type ServiceContext struct {
	Config config.Config
	//add dependency on user model
	UserModel model.UserModel
	//add dependency on user model
	UserReceiveAddressModel model.UserReceiveAddressModel
	UserCollectionModel     model.UserCollectionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                  c,
		UserModel:               model.NewUserModel(sqlConn, c.CacheRedis),
		UserReceiveAddressModel: model.NewUserReceiveAddressModel(sqlConn, c.CacheRedis),
		UserCollectionModel:     model.NewUserCollectionModel(sqlConn, c.CacheRedis),
	}
}
