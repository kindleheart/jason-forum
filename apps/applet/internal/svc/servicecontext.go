package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"jason-forum/apps/applet/internal/config"
	"jason-forum/apps/user/rpc/user"
	"jason-forum/pkg/interceptors"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRpc := zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	bizRedis, err := redis.NewRedis(c.BizRedis)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userRpc),
		BizRedis: bizRedis,
	}
}
