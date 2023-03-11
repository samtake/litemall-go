package middlewares

import (
	"github.com/gin-gonic/gin"
	"litemall-api/goods_web/models"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			/*暂不判断角色，改用Casbin*/
			//ctx.JSON(http.StatusForbidden, gin.H{
			//	"msg": "无权限",
			//})
			//ctx.Abort()
			//return
			/*暂不判断角色，改用Casbin*/
			ctx.Next()
			return
		}
		ctx.Next()
	}

}
