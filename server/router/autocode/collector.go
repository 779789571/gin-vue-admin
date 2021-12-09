package autocode

import (
	"github.com/779789571/gin-vue-admin/server/api/v1"
	"github.com/779789571/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CollectorRouter struct {
}

// InitCollectorRouter 初始化 Collector 路由信息
func (s *CollectorRouter) InitCollectorRouter(Router *gin.RouterGroup) {
	collectorRouter := Router.Group("collector").Use(middleware.OperationRecord())
	collectorRouterWithoutRecord := Router.Group("collector")
	var collectorApi = v1.ApiGroupApp.AutoCodeApiGroup.CollectorApi
	{
		collectorRouter.POST("createCollector", collectorApi.CreateCollector)             // 新建Collector
		collectorRouter.DELETE("deleteCollector", collectorApi.DeleteCollector)           // 删除Collector
		collectorRouter.DELETE("deleteCollectorByIds", collectorApi.DeleteCollectorByIds) // 批量删除Collector
		collectorRouter.PUT("updateCollector", collectorApi.UpdateCollector)              // 更新Collector
	}
	{
		collectorRouterWithoutRecord.GET("findCollector", collectorApi.FindCollector)       // 根据ID获取Collector
		collectorRouterWithoutRecord.GET("getCollectorList", collectorApi.GetCollectorList) // 获取Collector列表
	}
}
