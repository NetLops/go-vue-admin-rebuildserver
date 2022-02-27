package system

import "rebuildServer/config"

// System
// Description: 配置文件描述
type System struct {
	Config config.Server `json:"config"`
}
