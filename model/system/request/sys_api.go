package request

import (
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
)

// api分页条件查询及排序结构体
type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式: 生序false(默认)|降序true
}
