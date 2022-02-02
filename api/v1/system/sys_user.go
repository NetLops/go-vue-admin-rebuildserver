package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/common/response"
	"rebuildServer/model/system"
	systemReq "rebuildServer/model/system/request"
	"rebuildServer/utils"
)

// @Tag SysUser
// @Summary 用户登录
// @Produce application/json
// @Param data body systemReq.Login true "用户名，密码，验证码"
// @Success 200 {object} response.Response{data=systemRes.LoginResponse,msg=
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	_ = c.ShouldBindJSON(&l) // 写进结构里
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWIthMessage(err.Error(), c)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := userService.Login(u); err != nil {
			global.GVA_LOG.Error("登录失败！用户名不存在或者密码错误！", zap.Error(err))
			response.FailWIthMessage("用户名不存在或者密码错误", c)
		} else {
			//
		}
	}
}

func (b *BaseApi) tokenNext(c *gin.Context, user system.SysUser) {
	return
}
