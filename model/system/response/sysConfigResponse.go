package response

import "rebuildServer/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
