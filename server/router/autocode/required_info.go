package autocode

import (
	"github.com/779789571/gin-vue-admin/server/api/v1"
	"github.com/779789571/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RequiredInfoRouter struct {
}

// InitRequiredInfoRouter 初始化 RequiredInfo 路由信息
func (s *RequiredInfoRouter) InitrequiredInfoRouter(Router *gin.RouterGroup) {
	requiredInfoRouter := Router.Group("requiredInfo").Use(middleware.OperationRecord())
	requiredInfoRouterWithoutRecord := Router.Group("requiredInfo")
	var requiredInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.RequiredInfoApi
	{
		requiredInfoRouter.POST("createRequiredInfo", requiredInfoApi.CreaterequiredInfo)             // 新建requiredInfo
		requiredInfoRouter.DELETE("deleteRequiredInfo", requiredInfoApi.DeleterequiredInfo)           // 删除requiredInfo
		requiredInfoRouter.DELETE("deleteRequiredInfoByIds", requiredInfoApi.DeleterequiredInfoByIds) // 批量删除requiredInfo
		requiredInfoRouter.PUT("updateRequiredInfo", requiredInfoApi.UpdaterequiredInfo)              // 更新requiredInfo
	}
	{
		requiredInfoRouterWithoutRecord.GET("findRequiredInfo", requiredInfoApi.FindrequiredInfo)       // 根据ID获取requiredInfo
		requiredInfoRouterWithoutRecord.GET("getRequiredInfoList", requiredInfoApi.GetrequiredInfoList) // 获取requiredInfo列表
	}
}
