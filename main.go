package main

import (
	"rebuildServer/core"
	"rebuildServer/global"
	"rebuildServer/initialize"
)

func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.GVA_LOG = core.Zap()       // 初始化zap日志库
	global.GVA_DB = initialize.Gorm() // Gorm 连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束钱关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

}
