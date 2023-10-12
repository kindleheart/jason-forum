package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/singleflight"
	"jason-forum/apps/article/rpc/internal/config"
	"jason-forum/apps/article/rpc/internal/model"
)

type ServiceContext struct {
	Config            config.Config
	ArticleModel      model.ArticleModel
	BizRedis          *redis.Redis
	SingleFlightGroup singleflight.Group
}

func NewServiceContext(c config.Config) *ServiceContext {
	bizRedis, err := redis.NewRedis(c.BizRedis)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:       c,
		ArticleModel: model.NewArticleModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		BizRedis:     bizRedis,
	}
}
