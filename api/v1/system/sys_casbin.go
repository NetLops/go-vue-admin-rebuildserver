package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/common/response"
	"rebuildServer/model/system/request"
	systemRes "rebuildServer/model/system/response"
	"rebuildServer/utils"
)

type CasbinApi struct{}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id，权限模型列表"
// @Success 200 {object} response.Response{msg=string} "更新角色api权限"
// @Router /casbin/UpdateCasbin [post]
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInRecive
	_ = c.ShouldBindJSON(&cmr)
	if err := utils.Verify(cmr, utils.AuthorityIdVerify); err != nil {
		response.FailWIthMessage(err.Error(), c)
		return
	}
	if err := CasbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWIthMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Casbin
// @Summary 获取权限猎豹
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id， 权限模型列表"
// @Success 200 {object} response.Response{data=systemRes.PolicyPathResponse, msg=string} "获取权限列表，返回包括csbin详情列表"
// @Router /casbin/getpolicyPathByAuthorityId [post]
func (cas *CasbinApi) GetPolicyPathAuthorityId(c *gin.Context) {
	var casbin request.CasbinInRecive
	_ = c.ShouldBindJSON(&casbin)
	if err := utils.Verify(casbin, utils.AuthorityIdVerify); err != nil {
		response.FailWIthMessage(err.Error(), c)
		return
	}
	paths := CasbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}
