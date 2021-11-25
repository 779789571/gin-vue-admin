package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiTokenApi struct {
}

var apiTokenService = service.ServiceGroupApp.AutoCodeServiceGroup.ApiTokenService

// CreateApiToken 创建ApiToken
// @Tags ApiToken
// @Summary 创建ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiToken true "创建ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiToken/createApiToken [post]
func (apiTokenApi *ApiTokenApi) CreateApiToken(c *gin.Context) {
	var apiToken autocode.ApiToken
	_ = c.ShouldBindJSON(&apiToken)
	if err := apiTokenService.CreateApiToken(apiToken); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApiToken 删除ApiToken
// @Tags ApiToken
// @Summary 删除ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiToken true "删除ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiToken/deleteApiToken [delete]
func (apiTokenApi *ApiTokenApi) DeleteApiToken(c *gin.Context) {
	var apiToken autocode.ApiToken
	_ = c.ShouldBindJSON(&apiToken)
	if err := apiTokenService.DeleteApiToken(apiToken); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApiTokenByIds 批量删除ApiToken
// @Tags ApiToken
// @Summary 批量删除ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apiToken/deleteApiTokenByIds [delete]
func (apiTokenApi *ApiTokenApi) DeleteApiTokenByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := apiTokenService.DeleteApiTokenByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApiToken 更新ApiToken
// @Tags ApiToken
// @Summary 更新ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ApiToken true "更新ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiToken/updateApiToken [put]
func (apiTokenApi *ApiTokenApi) UpdateApiToken(c *gin.Context) {
	var apiToken autocode.ApiToken
	_ = c.ShouldBindJSON(&apiToken)
	if err := apiTokenService.UpdateApiToken(apiToken); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApiToken 用id查询ApiToken
// @Tags ApiToken
// @Summary 用id查询ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ApiToken true "用id查询ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiToken/findApiToken [get]
func (apiTokenApi *ApiTokenApi) FindApiToken(c *gin.Context) {
	var apiToken autocode.ApiToken
	_ = c.ShouldBindQuery(&apiToken)
	if err, reapiToken := apiTokenService.GetApiToken(apiToken.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapiToken": reapiToken}, c)
	}
}

// GetApiTokenList 分页获取ApiToken列表
// @Tags ApiToken
// @Summary 分页获取ApiToken列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ApiTokenSearch true "分页获取ApiToken列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiToken/getApiTokenList [get]
func (apiTokenApi *ApiTokenApi) GetApiTokenList(c *gin.Context) {
	var pageInfo autocodeReq.ApiTokenSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := apiTokenService.GetApiTokenInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
