package system

import (
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"rebuildServer/global"
	"rebuildServer/model/system"
)

var User = new(user)

type user struct {
}

func (u *user) TableName() string {
	return "sys_users"
}

func (u *user) Initialize() error {
	entities := []system.SysUser{
		{UUID: uuid.NewV4(), Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "超级管理员", HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{UUID: uuid.NewV4(), Username: "user123", Password: "e10adc3949ba59abbe56e057f20f883e", NickName: "QMPlusUser", HeaderImg: "https:///qmplusimg.henrongyi.top/1572075907logo.png", AuthorityId: "9528"},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, u.TableName()+"表数据初始化失败！")
	}
	return nil
}

func (u *user) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("username = ?", "user123").First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
