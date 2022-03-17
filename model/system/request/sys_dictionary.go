package request

import (
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
