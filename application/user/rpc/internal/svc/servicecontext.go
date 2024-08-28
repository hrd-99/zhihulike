package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zhihulike/application/user/rpc/internal/config"
	"zhihulike/application/user/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
