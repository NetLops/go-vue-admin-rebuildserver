package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/common/request"
	"rebuildServer/model/common/response"
	"rebuildServer/model/system"
	systemReq "rebuildServer/model/system/request"
	"rebuildServer/utils"
)

type OperationRecordApi struct {
}

// CreateSysOperationRecord
// @Tags SysOperationRecord
// @Summary 创建SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "创建SysOperationRecord"
// @Success 2000 {object} response.Response{msg=string} "创建SysOperationRecord"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func (s *OperationRecordApi) CreateSysOperationRecord(c *gin.Context) {
	var sysoperationRecord system.SysOperationRecord
	_ = c.ShouldBindJSON(&sysoperationRecord)
	if err := operationRecordService.CreateSysOperationRecord(sysoperationRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWIthMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysOperationRecord
// @Tags SysOperationRecord
// @Summary 删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysOperationRecord true "SysOperationRecord模型"
// @Success 200 {object} response.Response{msg=string} "删除SysOperationRecord"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *OperationRecordApi) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := operationRecordService.DeleteSysOperationRecord(sysOperationRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWIthMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysOperationRecordByIds
// @Tags SysOperationRecord
// @Summary 批量删除SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysOperationRecord"
// @Success 200 {object} response.Response{msg=string} "批量删除SysOperationRecord"
// @Router /sysOperationRecord/deleteSysOperationRecordIds [delete]
func (s *OperationRecordApi) DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := operationRecordService.DeleteSysOperationRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWIthMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindSysOperationRecord
// @Tags SysOperationRecord
// @Summary 用id查询SysOperationRecord
// @Security ApiKeyAuth
// @accept application/json
// Produce application/json
// @Param data query system.SysOperationRecord true "Id"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询SysOperationRecord"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func (s *OperationRecordApi) FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := utils.Verify(sysOperationRecord, utils.IdVerify); err != nil {
		response.FailWIthMessage(err.Error(), c)
		return
	}
	if err, resysOperationRecord := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWIthMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysOperationRecord": resysOperationRecord}, "查询成功", c)
	}
}

// GetSysOperationRecordList
// @Tags SysOperationRecord
// @Summary 分页获取SysOperationRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysOperationRecordSearch true "页码，每页大小，索索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取SysOperationRecord列表，返回包括列表，总数，页码，每页数量"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (s *OperationRecordApi) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo systemReq.SysOperationSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := operationRecordService.GetSysOperationRecordInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWIthMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
