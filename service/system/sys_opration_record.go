package system

import (
	"rebuildServer/global"
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
	systemReq "rebuildServer/model/system/request"
)

type OperationRecordService struct {
}

// CreateSysOperationRecord
// description: 创建记录
// param: sysOperationRecord model.SysOperationRecord
// return: err error
func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}

// DeleteSysOperationRecordByIds
//
// Description: 批量删除记录
//
// receiver: operationRecordService
//
// param: ids
//
// return: err
func (operationRecordService *OperationRecordService) DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

// DeleteSysOperationRecord
//
// Description: 删除操作记录
//
// receiver: operationRecordService
//
// param: sysOperationRecord
//
// return: err
func (operationRecordService *OperationRecordService) DeleteSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GVA_DB.Delete(&sysOperationRecord).Error
	return err
}

// GetSysOperationRecord
//
// Description: 根据id获取单条操作记录
//
// receiver: operationRecordService
//
// param: id uint
//
// return: err error
// return: sysOperationRecord system.SysOperationRecord
func (operationRecordService *OperationRecordService) GetSysOperationRecord(id uint) (err error, sysOperationRecord system.SysOperationRecord) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return
}

func (operationRecordService *OperationRecordService) GetSysOperationRecordInfoList(info systemReq.SysOperationSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysOperationRecord{})
	var sysOperationRecords []system.SysOperationRecord
	// 如果有条件搜索 下方会自动创建检索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return err, sysOperationRecords, total
}
