package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"jason-forum/apps/like/mq/internal/config"
	"jason-forum/apps/like/mq/internal/logic"
	"jason-forum/apps/like/mq/internal/svc"
)

var configFile = flag.String("f", "etc/like.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range logic.Consumers(ctx, svcCtx) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
}
