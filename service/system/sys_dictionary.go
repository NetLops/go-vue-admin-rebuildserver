package system

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"rebuildServer/global"
	"rebuildServer/model/system"
	"rebuildServer/model/system/request"
)

type DictionaryService struct {
}

// CreateSysDictionary
//
// Description: 创建字典数据
//
// receiver: dictionaryService
//
// param: sysDictionary system.SysDictionary
//
// return: err error
func (dictionaryService *DictionaryService) CreateSysDictionary(sysDictionary system.SysDictionary) (err error) {
	if (!errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GVA_DB.Create(&sysDictionary).Error
	return err
}

// DeleteSysDictionary
//
// Description: 删除字典数据
//
// receiver: dictionaryService
//
// param: sysDictionary system.SysDictionary
//
// return: err error
func (dictionaryService *DictionaryService) DeleteSysDictionary(sysDictionary system.SysDictionary) (err error) {
	err = global.GVA_DB.Where("id = ?", sysDictionary.ID).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&sysDictionary).Error
	if err != nil {
		return err
	}

	return global.GVA_DB.Model(&system.SysDictionaryDetail{}).Delete(sysDictionary.SysDictionaryDetails).Error
}

// UpdateSysDictionary
//
// Description: 更新字典数据
//
// receiver: dictionaryService
//
// param: sysDictionary system.SysDictionary
//
// return: err error
func (dictionaryService *DictionaryService) UpdateSysDictionary(sysDictionary *system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	db := global.GVA_DB.Where("id = ?", sysDictionary.ID).First(&dict)
	if dict.Type != sysDictionary.Type {
		if !errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许创建")
		}
	}
	err = db.Updates(sysDictionaryMap).Error
	return err
}

// GetSysDictionary
//
// Description: 根据id或者type获取字典单条数据
//
// receiver: dictionaryService
//
// param: Type string
// param: Id uint
//
// return: err error
// return: sysDictionary system.SysDictionary
func (dictionaryService *DictionaryService) GetSysDictionary(Type string, Id uint) (err error, sysDictionary system.SysDictionary) {
	err = global.GVA_DB.Where("type = ? OR id = ? and status = ?", Type, Id, true).Preload("SysDictionaryDetails", "status = ?", true).First(&sysDictionary).Error
	return
}

// GetSysDictionaryInfoList
//
// Description: 分页获取字典列表
//
// receiver: dictionaryService
//
// param: info request.SysDictionarySearch
//
// return: err error
// return: list interface{}
// return: total int64
func (dictionaryService *DictionaryService) GetSysDictionaryInfoList(info request.SysDictionarySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysDictionary{})
	var sysDictionarys []system.SysDictionary
	// 如果有条件检索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return err, sysDictionarys, total
}
