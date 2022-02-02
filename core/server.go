package core

import (
	"rebuildServer/global"
	"rebuildServer/initialize"
	"rebuildServer/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis 服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

}
