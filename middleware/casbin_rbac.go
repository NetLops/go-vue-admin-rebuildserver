package middleware

import (
	"github.com/gin-gonic/gin"
	"rebuildServer/global"
	"rebuildServer/model/common/response"
	"rebuildServer/service"
	"rebuildServer/utils"
)

var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService

// 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		waitUse, _ := utils.GetClaims(context)
		// 获取请求的PATH
		obj := context.Request.URL.Path
		// 获取请求方法
		act := context.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := casbinService.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.GVA_CONFIG.System.Env == "develop" || success {
			context.Next()
		} else {
			response.FailWithDetailed(gin.H{}, "权限不足", context)
			context.Abort()
			return
		}
	}
}
