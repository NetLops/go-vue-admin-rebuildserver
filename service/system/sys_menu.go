package system

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"rebuildServer/global"
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
	"strconv"
)

type MenuService struct {
}

var MenuServiceApp = new(MenuService)

// getMenuTreeMap
//
// Description: 获取路遇总树map
//
// receiver: menuService
//
// param: authorityId string
//
// return: err error
// return: treeMap map[string][]system.SysMenu
func (menuService *MenuService) getMenuTreeMap(authorityId string) (err error, treeMap map[string][]system.SysMenu) {
	var allMenus []system.SysMenu
	treeMap = make(map[string][]system.SysMenu)
	err = global.GVA_DB.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// getMenuTree
//
// Description: 获取动态菜单树
//
// receiver: menuService
//
// param: authorityId string
//
// return: err error
// return: menus []system.SysMenu
func (menuService *MenuService) GetMenuTree(authorityId string) (err error, menus []system.SysMenu) {
	err, menuTree := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

// GetChildrenList
//
// Description: 获取子菜单
//
// receiver: menuService
//
// param: menu *system.SysMenu
// param: treeMap map[string][]system.SysMenu
//
// return: err error
func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[string][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// GetIinfoList
//
// Description: 获取路由分页
//
// receiver: menuService
//
//
// return: err error
// return: list interface{}
// return: total int64
func (menuService *MenuService) GetInfoList() (err error, list interface{}, total int64) {
	var menuList []system.SysBaseMenu
	err, treeMap := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

// getBaseChildrenList
//
// Description: 获取菜单的子菜单
//
// receiver: menuService
//
// param: menu *system.SysBaseMenu
// param: treeMap map[string][]system.SysBaseMenu
//
// return: err error
func (menuService *MenuService) getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[string][]system.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// AddBaseMenu
//
// Description: 添加基础路由
//
// receiver: menuService
//
// param: menu system.SysBaseMenu
//
// return: error
func (menuService *MenuService) AddBaseMenu(menu system.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.GVA_DB.Create(&menu).Error
}

// getBaseMenuTreeMap
//
// Description: 获取路由器总树Map
//
// receiver: menuService
//
//
// return: err error
// return: treeMap map[string][]system.SysBaseMenu
func (menuService *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]system.SysBaseMenu) {
	var allMenus []system.SysBaseMenu
	treeMap = make(map[string][]system.SysBaseMenu)
	err = global.GVA_DB.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// GetBaseMenuTree
//
// Description: 获取基础路由树
//
// receiver: menuService
//
//
// return: err error
// return: menus []system.SysBaseMenu
func (menuService *MenuService) GetBaseMenuTree() (err error, menus []system.SysBaseMenu) {
	err, treeMap := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

// AddMenuAuthority
//
// Description: 为角色增加menu树
//
// receiver: menuService
//
// param: menus []system.SysBaseMenu
// param: authorityId string
//
// return: err error
func (menuService *MenuService) AddMenuAuthority(menus []system.SysBaseMenu, authorityId string) (err error) {
	var auth system.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []system.SysMenu) {
	err = global.GVA_DB.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authorityu_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authroity_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	//err = global.GVA_DB.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}
