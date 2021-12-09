package autocode

import (
	"github.com/779789571/gin-vue-admin/server/api/v1"
	"github.com/779789571/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RulesRouter struct {
}

// InitRulesRouter 初始化 Rules 路由信息
func (s *RulesRouter) InitRulesRouter(Router *gin.RouterGroup) {
	rulesRouter := Router.Group("rules").Use(middleware.OperationRecord())
	rulesRouterWithoutRecord := Router.Group("rules")
	var rulesApi = v1.ApiGroupApp.AutoCodeApiGroup.RulesApi
	{
		rulesRouter.POST("createRules", rulesApi.CreateRules)             // 新建Rules
		rulesRouter.DELETE("deleteRules", rulesApi.DeleteRules)           // 删除Rules
		rulesRouter.DELETE("deleteRulesByIds", rulesApi.DeleteRulesByIds) // 批量删除Rules
		rulesRouter.PUT("updateRules", rulesApi.UpdateRules)              // 更新Rules
	}
	{
		rulesRouterWithoutRecord.GET("findRules", rulesApi.FindRules)       // 根据ID获取Rules
		rulesRouterWithoutRecord.GET("getRulesList", rulesApi.GetRulesList) // 获取Rules列表
	}
}
