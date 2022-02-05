package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"rebuildServer/global"
	"rebuildServer/model/common/response"
	"rebuildServer/model/system/request"
)

type DBApi struct {
}

// InitDB
// @Tags InitDB
// @Summary 初始化用户数据库
func (i *DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("已存在数据库配置!")
		response.FailWIthMessage("已存在数据库配置", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("参数检验不通过!", zap.Error(err))
		response.FailWIthMessage("参数检验不通过", c)
	}

}
