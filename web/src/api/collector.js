import service from '@/utils/request'

// @Tags Collector
// @Summary 创建Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Collector true "创建Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /collector/createCollector [post]
export const createCollector = (data) => {
  return service({
    url: '/collector/createCollector',
    method: 'post',
    data
  })
}

// @Tags Collector
// @Summary 删除Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Collector true "删除Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /collector/deleteCollector [delete]
export const deleteCollector = (data) => {
  return service({
    url: '/collector/deleteCollector',
    method: 'delete',
    data
  })
}

// @Tags Collector
// @Summary 删除Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /collector/deleteCollector [delete]
export const deleteCollectorByIds = (data) => {
  return service({
    url: '/collector/deleteCollectorByIds',
    method: 'delete',
    data
  })
}

// @Tags Collector
// @Summary 更新Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Collector true "更新Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /collector/updateCollector [put]
export const updateCollector = (data) => {
  return service({
    url: '/collector/updateCollector',
    method: 'put',
    data
  })
}

// @Tags Collector
// @Summary 用id查询Collector
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Collector true "用id查询Collector"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /collector/findCollector [get]
export const findCollector = (params) => {
  return service({
    url: '/collector/findCollector',
    method: 'get',
    params
  })
}

// @Tags Collector
// @Summary 分页获取Collector列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Collector列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /collector/getCollectorList [get]
export const getCollectorList = (params) => {
  return service({
    url: '/collector/getCollectorList',
    method: 'get',
    params
  })
}
