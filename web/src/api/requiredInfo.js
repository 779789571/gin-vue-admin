import service from '@/utils/request'

// @Tags requiredInfo
// @Summary 创建requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.requiredInfo true "创建requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /requiredInfo/createrequiredInfo [post]
export const createrequiredInfo = (data) => {
  return service({
    url: '/requiredInfo/createRequiredInfo',
    method: 'post',
    data
  })
}

// @Tags requiredInfo
// @Summary 删除requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.requiredInfo true "删除requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /requiredInfo/deleterequiredInfo [delete]
export const deleterequiredInfo = (data) => {
  return service({
    url: '/requiredInfo/deleteRequiredInfo',
    method: 'delete',
    data
  })
}

// @Tags requiredInfo
// @Summary 删除requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /requiredInfo/deleterequiredInfo [delete]
export const deleterequiredInfoByIds = (data) => {
  return service({
    url: '/requiredInfo/deleteRequiredInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags requiredInfo
// @Summary 更新requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.requiredInfo true "更新requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /requiredInfo/updaterequiredInfo [put]
export const updaterequiredInfo = (data) => {
  return service({
    url: '/requiredInfo/updateRequiredInfo',
    method: 'put',
    data
  })
}

// @Tags requiredInfo
// @Summary 用id查询requiredInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.requiredInfo true "用id查询requiredInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /requiredInfo/findrequiredInfo [get]
export const findrequiredInfo = (params) => {
  return service({
    url: '/requiredInfo/findRequiredInfo',
    method: 'get',
    params
  })
}

// @Tags requiredInfo
// @Summary 分页获取requiredInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取requiredInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /requiredInfo/getrequiredInfoList [get]
export const getrequiredInfoList = (params) => {
  return service({
    url: '/requiredInfo/getRequiredInfoList',
    method: 'get',
    params
  })
}
