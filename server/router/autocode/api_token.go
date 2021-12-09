package autocode

import (
	"github.com/779789571/gin-vue-admin/server/api/v1"
	"github.com/779789571/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiTokenRouter struct {
}

// InitApiTokenRouter 初始化 ApiToken 路由信息
func (s *ApiTokenRouter) InitApiTokenRouter(Router *gin.RouterGroup) {
	apiTokenRouter := Router.Group("apiToken").Use(middleware.OperationRecord())
	apiTokenRouterWithoutRecord := Router.Group("apiToken")
	var apiTokenApi = v1.ApiGroupApp.AutoCodeApiGroup.ApiTokenApi
	{
		apiTokenRouter.POST("createApiToken", apiTokenApi.CreateApiToken)             // 新建ApiToken
		apiTokenRouter.DELETE("deleteApiToken", apiTokenApi.DeleteApiToken)           // 删除ApiToken
		apiTokenRouter.DELETE("deleteApiTokenByIds", apiTokenApi.DeleteApiTokenByIds) // 批量删除ApiToken
		apiTokenRouter.PUT("updateApiToken", apiTokenApi.UpdateApiToken)              // 更新ApiToken
	}
	{
		apiTokenRouterWithoutRecord.GET("findApiToken", apiTokenApi.FindApiToken)       // 根据ID获取ApiToken
		apiTokenRouterWithoutRecord.GET("getApiTokenList", apiTokenApi.GetApiTokenList) // 获取ApiToken列表
	}
}
