package v1

import (
	"rebuildServer/api/v1/autocode"
	"rebuildServer/api/v1/example"
	"rebuildServer/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
