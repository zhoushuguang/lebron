package svc

import (
	"github.com/zhoushuguang/lebron/apps/user/model"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//add a dependency on user model
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
