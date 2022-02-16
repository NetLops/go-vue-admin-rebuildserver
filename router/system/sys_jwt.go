package system

import (
	"github.com/gin-gonic/gin"
	v1 "rebuildServer/api/v1"
)

type JwtRouter struct {
}

func (s *JwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	jwtRouter := Router.Group("jwt")
	jwtApi := v1.ApiGroupApp.SystemApiGroup.JwtApi
	{
		jwtRouter.POST("jsonInBlackList", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
