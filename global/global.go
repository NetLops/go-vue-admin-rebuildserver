package global

import (
	"github.com/go-redis/redis"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"rebuildServer/config"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper

	GVA_LOG *zap.Logger
	//GVA_Timer
	GVA_Concurrency_Control = &singleflight.Group{}

	BlackCache local_cache.Cache

	lock sync.RWMutex
)

//
//  GetGlobalDBByDBName
//  @Description: 通过名称获取db List中的db
//  @param dbname
//  @return *gorm.DB
//
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

//
//  MustGetGlobalDBByDBName
//  @Description: 通过名称获取db，如果不存在则panic
//  @param dbname
//  @return *gorm.DB
//
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("do no init")
	}
	return db
}
