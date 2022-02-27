package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/common/response"
	"rebuildServer/model/system"
	systemRes "rebuildServer/model/system/response"
	"rebuildServer/utils"
)

type SystemApi struct {
}

// @Tags System
// @Summary 权限配置文件内容
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.SysConfigResponse,msg=string} "获取配置文件内容，返回包括系统配置"
// @Router /system/getSystemConfig [post]
func (s *SystemApi) GetSystemConfig(c *gin.Context) {
	if err, config := systemConfigService.GetSystemConfig(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWIthMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysConfigResponse{Config: config}, "获取成功", c)
	}
}

// @Tags System
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.System true "设置配置文件内容"
// @Success 200 {object} response.Response{data=string} "设置配置文件内容"
// @Router /system/setSystemConfig [post]
func (s *SystemApi) SetSystemConfig(c *gin.Context) {
	var sys system.System
	_ = c.ShouldBindJSON(&sys)
	if err := systemConfigService.SetSystemConfig(sys); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWIthMessage("设置失败", c)
	} else {
		response.OkWithData("设置成功", c)
	}
}

// @Tags System
// @Summary 重启系统
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "重启系统"
// @Router /system/reloadSystem [post]
func (s *SystemApi) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		global.GVA_LOG.Error("重启系统失败!", zap.Error(err))
		response.FailWIthMessage("重启系统失败", c)
	} else {
		response.OkWithMessage("重启系统成功", c)
	}
}

// @Tags System
// @Summary 获取服务器信息
// @Security ApiKeyAuth
// @Produce application/json
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取服务器信息"
// @Router /system/getServerInfo [post]
func (s *SystemApi) GetServerInfo(c *gin.Context) {
	if server, err := systemConfigService.GetServerInfo(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWIthMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}
