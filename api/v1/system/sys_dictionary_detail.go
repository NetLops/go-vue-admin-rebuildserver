package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/common/response"
	"rebuildServer/model/system"
	"rebuildServer/model/system/request"
	"rebuildServer/utils"
)

type DictionaryDetailApi struct {
}

// CreateSysDictionaryDetail
// @Tags SysDictionaryDetail
// @Summary 创建SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionaryDetail true "SysDictionaryDetail模型"
// @Success 200 {object} response.Response{msg=string} "创建SysDictionaryDetail"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func (s *DictionaryDetailApi) CreateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := dictionaryDetailService.CreateSysDictionaryDetail(detail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWIthMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysDictionaryDetail
// @Tags SysDictionaryDetail
// @Summary 更新SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce applocation/json
// @Param data body system.SysDictionaryDetail true "更新SysDictionaryDetail"
// @Success 200 {object} response.Response{msg=string} "更新SysDictionaryDetail"
// @Router /sysDictionary/updaetSysDictionaryDetail [put]
func (s *DictionaryDetailApi) DeleteSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := dictionaryDetailService.DeleteSysDictionaryDetail(detail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWIthMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateSysDictionaryDetail
// @Tags SysDictionaryDetail
// @Summary 更新SysDictionaryDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictioanryDetail true "更新SysDictionaryDetail"
// @Success 200 {object} response.Response{msg=string} "更新SysDictionaryDetail"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (s *DictionaryDetailApi) UpdateSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := dictionaryDetailService.UpdateSysDictionaryDetail(&detail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWIthMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysDictionaryDetail
// @Tags SysDictionaryDetail
// @Summary 用id查询SysDictionaryDetail
// @Security ApiKeyAuth
// @acept application/json
// @Produce application/json
// @param data query system.SysDictionaryDetail true "用id查询SysDictionaryDetail"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询SysDictionaryDetail"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func (s *DictionaryDetailApi) FindSysDictionaryDetail(c *gin.Context) {
	var detail system.SysDictionaryDetail
	_ = c.ShouldBindJSON(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWIthMessage(err.Error(), c)
		return
	}
	if err, resysDictionaryDetail := dictionaryDetailService.GetSysDictionaryDetail(detail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWIthMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictionaryDetail": resysDictionaryDetail}, "查询成功", c)
	}
}

// GetSysDictionaryDetailList
// @Tags SysDictionaryDetail
// @Summary 分页获取SysDictionaryDetail 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysDictionaryDetailSearch true "页码，每页大小，搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取SysDictionaryDetail列表,返回包括列表，总数，页码，每页数量"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (s *DictionaryDetailApi) GetSysDictionaryDetailList(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := dictionaryDetailService.GetSysDictionaryDetailInfoList(pageInfo); err != nil {
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
