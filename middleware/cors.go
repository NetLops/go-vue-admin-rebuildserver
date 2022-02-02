package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rebuildServer/config"
	"rebuildServer/global"
)

// Cors 直接放行所有跨域请求并放行所有 OPTIONS 方法
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		context.Header("Access-Control-Allow-Origin", origin)
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-Id")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE,PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		context.Next()
	}
}

// CorsByRules 按照配置处理跨域请求
func CorsByRules() gin.HandlerFunc {
	// 放行全部
	if global.GVA_CONFIG.Cors.Mode == "allow-all" {
		return Cors()
	}
	return func(context *gin.Context) {
		whilteList := checkCors(context.GetHeader("origin"))

		// 通过检查，添加请求头
		if whilteList != nil {
			context.Header("Access-Control-Allow-Origin", whilteList.AllowOrigin)
			context.Header("Access-Control-Allow-Headers", whilteList.AllowHeaders)
			context.Header("Access-Control-Allow-Methods", whilteList.AllowMethods)
			context.Header("Access-Control-Allow-Headers", whilteList.ExposeHeaders)
			if whilteList.AllowCredentials {
				context.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		// 严格白名单模式未用过检查，直接拒绝处理请求
		if whilteList == nil && global.GVA_CONFIG.Cors.Mode == "strict-whiteList" && !(context.Request.Method == "GET" && context.Request.URL.Path == "/health") {
			context.AbortWithStatus(http.StatusForbidden)
		} else {
			// 非严格白名单模式，无论是否通过检查均放行所有 OPTIONS 方法
			if context.Request.Method == "OPTIONS" {
				context.AbortWithStatus(http.StatusNoContent)
			}
		}

		// 处理请求
		context.Next()
	}
}

func checkCors(currentOrigin string) *config.CORSWhiteList {
	for _, whiteList := range global.GVA_CONFIG.Cors.WhiteList {
		// 遍历配置重的跨域头，寻找匹配项
		if currentOrigin == whiteList.AllowOrigin {
			return &whiteList
		}
	}
	return nil
}
