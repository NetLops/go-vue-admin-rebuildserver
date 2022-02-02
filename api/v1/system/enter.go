package system

import "rebuildServer/service"

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
