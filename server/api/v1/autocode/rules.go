package autocode

import (
	"github.com/779789571/gin-vue-admin/server/global"
	"github.com/779789571/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/779789571/gin-vue-admin/server/model/autocode/request"
	"github.com/779789571/gin-vue-admin/server/model/common/request"
	"github.com/779789571/gin-vue-admin/server/model/common/response"
	"github.com/779789571/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RulesApi struct {
}

var rulesService = service.ServiceGroupApp.AutoCodeServiceGroup.RulesService

// CreateRules 创建Rules
// @Tags Rules
// @Summary 创建Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Rules true "创建Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rules/createRules [post]
func (rulesApi *RulesApi) CreateRules(c *gin.Context) {
	var rules autocode.Rules
	_ = c.ShouldBindJSON(&rules)
	if err := rulesService.CreateRules(rules); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRules 删除Rules
// @Tags Rules
// @Summary 删除Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Rules true "删除Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rules/deleteRules [delete]
func (rulesApi *RulesApi) DeleteRules(c *gin.Context) {
	var rules autocode.Rules
	_ = c.ShouldBindJSON(&rules)
	if err := rulesService.DeleteRules(rules); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRulesByIds 批量删除Rules
// @Tags Rules
// @Summary 批量删除Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rules/deleteRulesByIds [delete]
func (rulesApi *RulesApi) DeleteRulesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := rulesService.DeleteRulesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRules 更新Rules
// @Tags Rules
// @Summary 更新Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Rules true "更新Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rules/updateRules [put]
func (rulesApi *RulesApi) UpdateRules(c *gin.Context) {
	var rules autocode.Rules
	_ = c.ShouldBindJSON(&rules)
	if err := rulesService.UpdateRules(rules); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRules 用id查询Rules
// @Tags Rules
// @Summary 用id查询Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Rules true "用id查询Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rules/findRules [get]
func (rulesApi *RulesApi) FindRules(c *gin.Context) {
	var rules autocode.Rules
	_ = c.ShouldBindQuery(&rules)
	if err, rerules := rulesService.GetRules(rules.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerules": rerules}, c)
	}
}

// GetRulesList 分页获取Rules列表
// @Tags Rules
// @Summary 分页获取Rules列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.RulesSearch true "分页获取Rules列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rules/getRulesList [get]
func (rulesApi *RulesApi) GetRulesList(c *gin.Context) {
	var pageInfo autocodeReq.RulesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := rulesService.GetRulesInfoList(pageInfo); err != nil {
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
