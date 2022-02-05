package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/common/response"
	"rebuildServer/model/system"
)

type JwtApi struct {
}

// @Tags Jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accpet application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "jwt加入黑名单"
// @Router /jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{
		Jwt: token,
	}
	if err := jwtService.JsonInBlacklist(jwt); err != nil {
		global.GVA_LOG.Error("jwt作废失败！", zap.Error(err))
		response.FailWIthMessage("jwt作废失败", c)
	} else {
		response.OkWithMessage("jwt作废成功", c)
	}

}
