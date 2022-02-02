package system

import (
	"github.com/gin-gonic/gin"
	"rebuildServer/middleware"
)

type ApiRouter struct {
}

// InitApiRouter 未搞完
func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())

}
