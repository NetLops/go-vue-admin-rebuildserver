package system

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"rebuildServer/global"
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
	"rebuildServer/model/system/response"
	"strconv"
)

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

// CreateAuthority
//
// Description: 创建一个角色
//
// receiver: authorityService
//
// param: auth system.SysAuthority
//
// return: err error
// return: authority system.SysAuthority
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (err error, authority system.SysAuthority) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.GVA_DB.Create(&auth).Error
	return err, auth
}

// CopyAuthority
//
// Description: 创建一个角色
//
// receiver: authorityService
//
// param: copyInfo response.SysAuthorityCopyResponse
//
// return: err error
// return: authority system.SysAuthority
func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority system.SysAuthority) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
	err, menus := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.GVA_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityName, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

// UpdateAuthority
//
// Description: 更改一个角色
//
// receiver: authorityService
//
// param: auth system.SysAuthority
//
// return: err error
// return: authority system.SysAuthority
func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (err error, authority system.SysAuthority) {
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

// DeleteAuthority
//
// Description: 删除角色
//
// receiver: authorityService
//
// param: auth *system.SysAuthority
//
// return: err error
func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) (err error) {
	if errors.Is(global.GVA_DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("parent_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.GVA_DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.DataAuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.SysBaseMenus) > 0 {
		err = global.GVA_DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
	} else {
		err = db.Error
		if err != nil {
			return
		}
	}
	err = global.GVA_DB.Delete(&[]system.SysAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error
	CasbinServiceApp.ClearCasbin(0, auth.AuthorityId)
	return err
}

// GetAuthorityInfoList
//
// Description: 分页获取数据
//
// receiver: authorityService
//
// param: info request.PageInfo
//
// return: err error
// return: list interface{}
// return: total int64
func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAuthority{})
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("Parent_id = ?", "0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = authorityService.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (err error, sa system.SysAuthority) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("pathMaps_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

// SetDataAuthority
//
// Description: 设置角色资源权限
//
// receiver: authorityService
//
// param: auth system.SysAuthority
//
// return: error
func (authorityService *AuthorityService) SetDataAuthority(auth system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("DataAuthority").Replace(&auth.DataAuthorityId)
	return err
}

// SetMenuAuthority
//
// Description: 菜单与角色绑定
//
// receiver: authorityService
//
// param: auth *system.SysAuthority
//
// return: error
func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// findChildrenAuthority
//
// Description: 查询子角色
//
// receiver: authorityService
//
// param: authority *system.SysAuthority
//
// return: err error
func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
