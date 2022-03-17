package request

import (
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
