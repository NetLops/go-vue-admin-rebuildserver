package system

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"rebuildServer/global"
	"rebuildServer/model/system"
	"rebuildServer/model/system/request"
	"sync"
)

type CasbinService struct {
}

var CasbinServiceApp = new(CasbinService)

// UpdateCasbin
//
// Description: 更新casbin权限
//
// receiver: casbinService
//
// param: authorityid string
// param: casbinInfos []request.CasbinInfo
//
// return: error
func (casbinService *CasbinService) UpdateCasbin(authorityid string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, authorityid)
	rules := [][]string{}
	for _, v := range casbinInfos {
		cm := system.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityid,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api，添加失败，请联系管理员")
	}
	return nil
}

// updateCasbinApi
//
// Description: API更新随动
//
// receiver: casbinService
//
// param: oldPath string
// param: newPath string
// param: oldMethod string
// param: newMethod string
//
// return: error
func (casbinService *CasbinService) updateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.GVA_DB.Table("casbin_rule").Model(&system.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldMethod, newMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// GetPolicyPathByAuthorityId
//
// Description: 获取权限列表
//
// receiver: casbinService
//
// param: authorityId string
//
// return: pathMaps []request.CasbinInfo
func (casbinService *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// ClearCasbin
//
// Description: 清除匹配的权限
//
// receiver: casbinService
//
// param: v int
// param: p ...string
//
// return: bool
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

// Casbin
//
// Description: 持久化到数据库 引入自定义规则
//
// receiver: casbinService
//
//
// return: *casbin.SyncedEnforcer
func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.GVA_DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.GVA_CONFIG.Casbin.ModelPath, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
