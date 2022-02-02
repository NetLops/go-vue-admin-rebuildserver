package request

import (
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
)

type SysOperationSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
