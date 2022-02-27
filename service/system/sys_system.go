package system

import (
	"go.uber.org/zap"
	"rebuildServer/config"
	"rebuildServer/global"
	"rebuildServer/model/system"
	"rebuildServer/utils"
)

type SystemConfigService struct {
}

// GetSystemConfig
//
// Description: 读取配置文件
//
// receiver: systemCOnfigService
//
//
// return: err error
// return: conf config.Server
func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.GVA_CONFIG
}

// SetSystemConfig
//
// Description: 设置配置文件
//
// receiver: systemCOnfigService
//
// param: system system.System
//
// return: err error
func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

// GetServerInfo
//
// Description: 获取服务器信息
//
// receiver: systemCOnfigService
//
//
// return: server *utils.Server
// return: err error
func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils,InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	return &s, nil
}
