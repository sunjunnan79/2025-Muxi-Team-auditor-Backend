// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
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

// Injectors from wire.go:

func InitWebServer(confPath string) *App {
	vipperSetting := viperx.NewVipperSetting(confPath)
	oAuthConfig := config.NewOAuthConf(vipperSetting)
	oAuthClient := client.NewOAuthClient(oAuthConfig)
	dbConfig := config.NewDBConf(vipperSetting)
	logConfig := config.NewLogConf(vipperSetting)
	logger := ioc.InitLogger(logConfig)
	db := ioc.InitDB(dbConfig, logger)
	userDAO := dao.NewUserDAO(db)
	authService := service.NewAuthService(userDAO)
	authController := controller.NewOAuthController(oAuthClient, authService)
	cacheConfig := config.NewCacheConf(vipperSetting)
	redisClient := ioc.InitCache(cacheConfig)
	jwtConfig := config.NewJWTConf(vipperSetting)
	redisJWTHandler := jwt.NewRedisJWTHandler(redisClient, jwtConfig)
	authMiddleware := middleware.NewAuthMiddleware(redisJWTHandler)
	corsMiddleware := middleware.NewCorsMiddleware()
	engine := router.NewRouter(authController, authMiddleware, corsMiddleware)
	appConf := config.NewAppConf(vipperSetting)
	app := NewApp(engine, appConf)
	return app
}