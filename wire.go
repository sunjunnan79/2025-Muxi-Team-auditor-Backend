//go:generate wire
//go:build wireinject

package main

import (
	"github.com/google/wire"
	"muxi_auditor/client"
	"muxi_auditor/config"
	"muxi_auditor/controller"
	"muxi_auditor/ioc"
	"muxi_auditor/middleware"
	"muxi_auditor/pkg/jwt"
	"muxi_auditor/pkg/viperx"
	"muxi_auditor/repository/dao"
	"muxi_auditor/router"
	"muxi_auditor/service"
)

func InitWebServer(confPath string) *App {
	wire.Build(
		viperx.NewVipperSetting,
		config.NewAppConf,
		config.NewJWTConf,
		config.NewOAuthConf,
		config.NewDBConf,
		config.NewLogConf,
		config.NewCacheConf,

		// 初始化基础依赖
		ioc.InitDB,
		ioc.InitLogger,
		ioc.InitCache,

		// 初始化具体模块
		dao.NewUserDAO,

		jwt.NewRedisJWTHandler,
		service.NewAuthService,
		controller.NewOAuthController,
		client.NewOAuthClient,
		router.NewRouter,

		// 中间件
		middleware.NewAuthMiddleware,
		middleware.NewCorsMiddleware,

		// 应用入口
		NewApp,
	)
	return &App{}
}