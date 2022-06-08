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

type CollectorApi struct {
}

var collectorService = service.ServiceGroupApp.AutoCodeServiceGroup.CollectorService

// CreateCollector 创建Collector
// @Tags Collector
// @Summary 创建Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Collector true "创建Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /collector/createCollector [post]
func (collectorApi *CollectorApi) CreateCollector(c *gin.Context) {
	var collector autocode.Collector
	_ = c.ShouldBindJSON(&collector)
	if err := collectorService.CreateCollector(collector); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCollector 删除Collector
// @Tags Collector
// @Summary 删除Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Collector true "删除Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /collector/deleteCollector [delete]
func (collectorApi *CollectorApi) DeleteCollector(c *gin.Context) {
	var collector autocode.Collector
	_ = c.ShouldBindJSON(&collector)
	if err := collectorService.DeleteCollector(collector); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCollectorByIds 批量删除Collector
// @Tags Collector
// @Summary 批量删除Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /collector/deleteCollectorByIds [delete]
func (collectorApi *CollectorApi) DeleteCollectorByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := collectorService.DeleteCollectorByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCollector 更新Collector
// @Tags Collector
// @Summary 更新Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Collector true "更新Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /collector/updateCollector [put]
func (collectorApi *CollectorApi) UpdateCollector(c *gin.Context) {
	var collector autocode.Collector
	_ = c.ShouldBindJSON(&collector)
	if err := collectorService.UpdateCollector(collector); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCollector 用id查询Collector
// @Tags Collector
// @Summary 用id查询Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Collector true "用id查询Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /collector/findCollector [get]
func (collectorApi *CollectorApi) FindCollector(c *gin.Context) {
	var collector autocode.Collector
	_ = c.ShouldBindQuery(&collector)
	if err, recollector := collectorService.GetCollector(collector.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recollector": recollector}, c)
	}
}

// GetCollectorList 分页获取Collector列表
// @Tags Collector
// @Summary 分页获取Collector列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CollectorSearch true "分页获取Collector列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /collector/getCollectorList [get]
func (collectorApi *CollectorApi) GetCollectorList(c *gin.Context) {
	var pageInfo autocodeReq.CollectorSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := collectorService.GetCollectorInfoList(pageInfo); err != nil {
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
