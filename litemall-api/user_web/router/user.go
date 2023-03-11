package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"litemall-api/user_web/api"
	middlewares2 "litemall-api/user_web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	zap.S().Infof("router 用户相关url")
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares2.JWTAuth(), middlewares2.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
	}

}
