package system

import (
	"rebuildServer/global"
	"rebuildServer/model/system"
	"rebuildServer/model/system/request"
)

type DictionaryDetailService struct {
}

// CreateSysDictionaryDetail
//
// Description: 创建字典详情数据
//
// receiver: dictionaryDetailService
//
// param: sysDictionaryDetail system.SysDictionaryDetail
//
// return: err error
func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

// DeleteSysDictionaryDetail
//
// Description: 删除字典详情数据
//
// receiver: dictionaryDetailService
//
// param: sysDictionaryDetail system.SysDictionaryDetail
//
// return: err error
func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Delete(&sysDictionaryDetail).Error
	return err
}

// UpdateSysDictionaryDetail
//
// Description: 更新字典详情数据
//
// receiver: dictionaryDetailService
//
// param: sysDictionaryDetail *system.SysDictionaryDetail
//
// return: err error
func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail *system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Save(sysDictionaryDetail).Error
	return err
}

// GetSysDictionaryDetail
//
// Description: 根据id获取字典详情单条数据
//
// receiver: dictionaryDetailService
//
// param: id uint
//
// return: err error
// return: sysDictionaryDetail system.SysDictionaryDetail
func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (err error, sysDictionaryDetail system.SysDictionaryDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

// GetSysDictionaryDetailInfoList
//
// Description: 分页获取字典详情列表
//
// receiver: dictionaryDetailService
//
// param: info request.SysDictionaryDetailSearch
//
// return: err error
// return: list interface{}
// return: total int64
func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {

	}
	err = db.Limit(limit).Offset(offset).Find(&sysDictionaryDetails).Error
	return err, sysDictionaryDetails, total
}
