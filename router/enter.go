package router

import (
	"rebuildServer/router/autocode"
	"rebuildServer/router/example"
	"rebuildServer/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
