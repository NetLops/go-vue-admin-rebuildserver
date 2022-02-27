package system

import "rebuildServer/service"

type ApiGroup struct {
	DBApi
	BaseApi
	SystemApi
	JwtApi
	SystemApiApi
	AuthorityMenuApi
	CasbinApi
	AutoCodeApi
}

var (
	apiService          = service.ServiceGroupApp.SystemServiceGroup.ApiService
	userService         = service.ServiceGroupApp.SystemServiceGroup.UserService
	jwtService          = service.ServiceGroupApp.SystemServiceGroup.JwtService
	initDBService       = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	menuService         = service.ServiceGroupApp.SystemServiceGroup.MenuService
	baseMenuService     = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	CasbinService       = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	autoCodeService     = service.ServiceGroupApp.SystemServiceGroup.AutoCodeService
)
