package router

import (
	"github.com/gin-gonic/gin"
	"muxi_auditor/api/request"
	"muxi_auditor/api/response"
	"muxi_auditor/pkg/ginx"
)

type OAuthController interface {
	Login(g *gin.Context, req request.LoginReq) (response.Response, error)
	Register(g *gin.Context, req request.RegisterReq) (response.Response, error)
	Logout(g *gin.Context) (response.Response, error)
	UpdateMyInfo(g *gin.Context, req request.UpdateUserReq) (response.Response, error)
	GetQiToken(g *gin.Context) (response.Response, error)
}

func RegisterOAuthRoutes(
	s *gin.RouterGroup,
	authMiddleware gin.HandlerFunc,
	c OAuthController,
) {
	//认证服务
	authGroup := s.Group("/user")
	authGroup.POST("/login", ginx.WrapReq(c.Login))
	authGroup.POST("/register", ginx.WrapReq(c.Register))
	authGroup.GET("/logout", authMiddleware, ginx.Wrap(c.Logout))
	authGroup.POST("/updateMyInfo", authMiddleware, ginx.WrapReq(c.UpdateMyInfo))
	authGroup.GET("/GetQiToken",authMiddleware,ginx.Wrap(c.GetQiToken))
}
