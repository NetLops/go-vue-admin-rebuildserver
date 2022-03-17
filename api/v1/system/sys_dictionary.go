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

type DictionaryApi struct {
}

// CreateSysDictionary
// @Tags SysDictionary
// @Summary 创建SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "创建SysDictionary"
// @Router /sysDictionary/createSysDictionary [post]
func (s *DictionaryApi) CreateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.CreateSysDictionary(dictionary); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWIthMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysDictionary
// @Tags SysDictionary
// @Summary 删除SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "删除sysDictionary"
// @Router /SysDictionary/deleteSysDictionary [delete]
func (s *DictionaryApi) DeleteSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.DeleteSysDictionary(dictionary); err != nil {
		global.GVA_LOG.Error("删除失败！", zap.Error(err))
		response.FailWIthMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// UpdateSysDictionary
// @Tags SysDictionary
// @Summary 更新SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictionary true "SysDictionary模型"
// @Success 200 {object} response.Response{msg=string} "更新SysDictionary"
// @Router /SysDictionary/updateSysDictionary [put]
func (s *DictionaryApi) UpdateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.UpdateSysDictionary(&dictionary); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWIthMessage("更新失败", c)
	}
}

// FindSysDictionary
// @Tags SysDictionary
// @Summary 用id查询SysDictionary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDictionary true "ID或字典名"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "用id查询SysDictionary"
// @Router /sysDictionary/findSysDictionary [get]
func (s *DictionaryApi) FindSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	_ = c.ShouldBindJSON(&dictionary)
	if err, sysDictionary := dictionaryService.GetSysDictionary(dictionary.Type, dictionary.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWIthMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "查询成功", c)
	}
}

// GetSysDictionaryList
// @Tags SysDictionary
// @Summary 分页获取SysDictionary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Parm data query request.SysDictionarySearch true "页码，每页大小，搜索条件"
// @Success 200 {object} response.Response{data=response.PageResult, msg=string} "分页获取SysDictionary列表，返回包括列表，总数、页码、每页数量"
// @Router /sysDictionary/getSysDictionaryList [get]
func (s *DictionaryApi) GetSysDictionaryList(c *gin.Context) {
	var pageInfo request.SysDictionarySearch
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWIthMessage(err.Error(), c)
		return
	}
	if err, list, total := dictionaryService.GetSysDictionaryInfoList(pageInfo); err != nil {
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
