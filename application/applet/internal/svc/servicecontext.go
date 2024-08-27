package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"zhihulike/application/applet/internal/config"
	"zhihulike/application/user/rpc/user"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//UserRPC:  user.NewUser(userRPC),
		BizRedis: redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
