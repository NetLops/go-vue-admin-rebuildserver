package system

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"rebuildServer/global"
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
	"rebuildServer/utils"
)

type UserService struct {
}

// Register
//
// Description: 用户注册
//
// receiver: userService
//
// param: u system.SysUser
//
// return: err error
// return: userInter system.SysUser
func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	var user system.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 用户注册
		return errors.New("用户名已注册"), userInter
	}

	// 否则 附加uuid 密码md5简单加密， 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

// Login
//
// Description: 用户登录
//
// receiver: userService
//
// param: u *system.SysUser
//
// return: err error
// return: userInter *system.SysUser
func (userService *UserService) Login(u *system.SysUser) (err error, userInter *system.SysUser) {
	if nil == global.GVA_DB {
		return fmt.Errorf("db not init"), nil
	}

	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authorities").Preload("Authority").First(&user).Error
	return err, &user
}

// ChangePassord
//
// Description: 修改用户密码
//
// receiver: userService
//
// param: u *system.SysUser
// param: newPassword string
//
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (err error, userInter *system.SysUser) {
	var user system.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// GetUserInfoList
//
// Description: 分月获取数据
//
// receiver: userService
//
// param: info request.PageInfo
//
// return: err error
// return: list interface{}
// return: total int64
func (userService *UserService) GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return err, userList, total
}

// SetUserAuthority
//
// Description: 设置一个用户的权限
//
// receiver: userService
//
// param: id uint
// param: uuid uuid.UUID
// param: authorityId string
//
// return: err error
func (userService *UserService) SetUserAuthority(id uint, uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUseAuthority{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}

	err = global.GVA_DB.Where("uuid = ?", uuid).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// SetUserAuthorities
//
// Description: 设置一个用户的权限
//
// receiver: userService
//
// param: id uint
// param: authorityIds []string
//
// return: err error
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []string) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		useAuthority := []system.SysUseAuthority{}
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, system.SysUseAuthority{
				id, v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

// DeleteUser
//
// Description: 删除用户
//
// receiver: userService
//
// param: id float64
//
// return: err error
func (userService *UserService) DeleteUser(id float64) (err error) {
	var user system.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return nil
	}
	err = global.GVA_DB.Delete(&[]system.SysUseAuthority{}, "sys_user_id = ?", id).Error
	return err
}

// SetUserInfo
//
// Description: 设置用户信息
//
// receiver: userService
//
// param: reqUser system.SysUser
//
// return: err error
// return: user system.SysUser
func (userService *UserService) SetUserInfo(reqUser system.SysUser) (err error, user system.SysUser) {
	err = global.GVA_DB.Updates(&reqUser).Error
	return err, reqUser
}

// GetUserInfo
//
// Description: 获取用户信息
//
// receiver: userService
//
// param: uuid uuid.UUID
//
// return: err error
// return: user system.SysUser
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (err error, user system.SysUser) {
	var reqUser system.SysUser
	err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	return err, reqUser
}

// FindUserById
//
// Description: 通过Id获取用户信息
//
// receiver: userService
//
// param: id int
//
// return: err error
// return: user *system.SysUser
func (userService *UserService) FindUserById(id int) (err error, user *system.SysUser) {
	var u system.SysUser
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// FindUserByUuid
//
// Description: 通过UUID获取用户信息
//
// receiver: userService
//
// param: uuid string
//
// return: err error
// return: user *system.SysUser
func (userService *UserService) FindUserByUuid(uuid string) (err error, user *system.SysUser) {
	var u system.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

// ResetPassword
//
// Description: 修改用户密码
//
// receiver: userService
//
// param: ID int
//
// return: err error
func (userService *UserService) ResetPassword(ID int) (err error) {
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.MD5V([]byte("123456"))).Error
	return err
}
