package system

import "rebuildServer/global"

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"comment:api路径"`           //	api 路径
	Description string `json:"description" gorm:"comment:apiZ中文描述"` // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`        // api组
	Method      string `json:"method" gorm:"comment:方法"`            // 方法：创建POST(默认)/查看GET/更新PUT/删除DELETE
}
