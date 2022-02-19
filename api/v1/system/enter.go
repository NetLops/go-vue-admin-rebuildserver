package system

import "rebuildServer/service"

type ApiGroup struct {
	DBApi
	BaseApi
	JwtApi
	SystemApiApi
	AuthorityMenuApi
}

var (
	apiService      = service.ServiceGroupApp.SystemServiceGroup.ApiService
	userService     = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService      = service.ServiceGroupApp.SystemServiceGroup.JwtService
	initDBService   = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	menuService     = service.ServiceGroupApp.SystemServiceGroup.MenuService
	baseMenuService = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
)
