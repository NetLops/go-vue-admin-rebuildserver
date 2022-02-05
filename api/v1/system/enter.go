package system

import "rebuildServer/service"

type ApiGroup struct {
	DBApi
	BaseApi
	JwtApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
	initService = service.ServiceGroupApp.SystemServiceGroup.InitService
)
