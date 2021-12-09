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

type RequiredInfoApi struct {
}

var requiredInfoService = service.ServiceGroupApp.AutoCodeServiceGroup.RequiredInfoService

// CreaterequiredInfo 创建requiredInfo
// @Tags requiredInfo
// @Summary 创建requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.RequiredInfo true "创建requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /requiredInfo/createrequiredInfo [post]
func (requiredInfoApi *RequiredInfoApi) CreaterequiredInfo(c *gin.Context) {
	var requiredInfo autocode.RequiredInfo
	_ = c.ShouldBindJSON(&requiredInfo)
	if err := requiredInfoService.CreaterequiredInfo(requiredInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleterequiredInfo 删除requiredInfo
// @Tags requiredInfo
// @Summary 删除requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.RequiredInfo true "删除requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /requiredInfo/deleterequiredInfo [delete]
func (requiredInfoApi *RequiredInfoApi) DeleterequiredInfo(c *gin.Context) {
	var requiredInfo autocode.RequiredInfo
	_ = c.ShouldBindJSON(&requiredInfo)
	if err := requiredInfoService.DeleterequiredInfo(requiredInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleterequiredInfoByIds 批量删除requiredInfo
// @Tags requiredInfo
// @Summary 批量删除requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /requiredInfo/deleterequiredInfoByIds [delete]
func (requiredInfoApi *RequiredInfoApi) DeleterequiredInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := requiredInfoService.DeleterequiredInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdaterequiredInfo 更新requiredInfo
// @Tags requiredInfo
// @Summary 更新requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.RequiredInfo true "更新requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /requiredInfo/updaterequiredInfo [put]
func (requiredInfoApi *RequiredInfoApi) UpdaterequiredInfo(c *gin.Context) {
	var requiredInfo autocode.RequiredInfo
	_ = c.ShouldBindJSON(&requiredInfo)
	if err := requiredInfoService.UpdaterequiredInfo(requiredInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindrequiredInfo 用id查询requiredInfo
// @Tags requiredInfo
// @Summary 用id查询requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.RequiredInfo true "用id查询requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /requiredInfo/findrequiredInfo [get]
func (requiredInfoApi *RequiredInfoApi) FindrequiredInfo(c *gin.Context) {
	var requiredInfo autocode.RequiredInfo
	_ = c.ShouldBindQuery(&requiredInfo)
	if err, rerequiredInfo := requiredInfoService.GetrequiredInfo(requiredInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerequiredInfo": rerequiredInfo}, c)
	}
}

// GetrequiredInfoList 分页获取requiredInfo列表
// @Tags requiredInfo
// @Summary 分页获取requiredInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.RequiredInfoSearch true "分页获取requiredInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /requiredInfo/getrequiredInfoList [get]
func (requiredInfoApi *RequiredInfoApi) GetrequiredInfoList(c *gin.Context) {
	var pageInfo autocodeReq.RequiredInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := requiredInfoService.GetrequiredInfoInfoList(pageInfo); err != nil {
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
