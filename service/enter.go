package service

import (
	"rebuildServer/service/autocode"
	"rebuildServer/service/example"
	"rebuildServer/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
