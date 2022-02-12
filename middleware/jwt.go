package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/common/response"
	"rebuildServer/model/system"
	"rebuildServer/service"
	"rebuildServer/utils"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登陆返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := context.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", context)
			context.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的账户异地登陆或令牌失效", context)
			context.Abort()
			return
		}
		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", context)
				context.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), context)
			context.Abort()
			return
		}
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开
		/*这部分没有写*/
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			context.Header("new-token", newToken)
			context.Header("new-expires-at", newToken)
			if global.GVA_CONFIG.System.UseMultipoint {
				err, redisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: redisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		context.Set("claims", claims)
		context.Next()
	}
}
