package system

import (
	"rebuildServer/global"
	"rebuildServer/model/common/request"
	"strconv"
	"strings"
)

// SysAutoCodeHistory 自动迁移代码记录，用于回滚，重放使用
type SysAutoCodeHistory struct {
	global.GVA_MODEL
	TableName     string `json:"tableName"`
	RequestMeta   string `gorm:"type:text" json:"requestMeta,omitempty"`   // 前端传入的结构化信息
	AutoCodePath  string `gorm:"type:text" json:"autoCodePath,omitempty"`  // 其他meta信息 path;path
	InjectionMeta string `gorm:"type:text" json:"injectionMeta,omitempty"` // 注入的内容RouterPath@functionName@RouterString
	StructName    string `json:"StructName"`
	StructCNName  string `json:"StructCNName"`
	ApiIDs        string `json:"apiIDs,omitempty"` // api表注册内容
	Flag          int    `json:"flag"`             // 标识对应状态0 代表创建，1 代表回滚 ...
}

// ToRequestIds ApiIDs 转换 request.IdsReq
func (m *SysAutoCodeHistory) ToRequestIds() request.IdsReq {
	if m.ApiIDs == "" {
		return request.IdsReq{}
	}
	slice := strings.Split(m.ApiIDs, ";")
	ids := make([]int, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		id, _ := strconv.ParseInt(slice[i], 10, 32)
		ids = append(ids, int(id))
	}
	return request.IdsReq{Ids: ids}
}
