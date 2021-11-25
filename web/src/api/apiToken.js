import service from '@/utils/request'

// @Tags ApiToken
// @Summary 创建ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiToken true "创建ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiToken/createApiToken [post]
export const createApiToken = (data) => {
  return service({
    url: '/apiToken/createApiToken',
    method: 'post',
    data
  })
}

// @Tags ApiToken
// @Summary 删除ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiToken true "删除ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiToken/deleteApiToken [delete]
export const deleteApiToken = (data) => {
  return service({
    url: '/apiToken/deleteApiToken',
    method: 'delete',
    data
  })
}

// @Tags ApiToken
// @Summary 删除ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiToken/deleteApiToken [delete]
export const deleteApiTokenByIds = (data) => {
  return service({
    url: '/apiToken/deleteApiTokenByIds',
    method: 'delete',
    data
  })
}

// @Tags ApiToken
// @Summary 更新ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ApiToken true "更新ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiToken/updateApiToken [put]
export const updateApiToken = (data) => {
  return service({
    url: '/apiToken/updateApiToken',
    method: 'put',
    data
  })
}

// @Tags ApiToken
// @Summary 用id查询ApiToken
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ApiToken true "用id查询ApiToken"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiToken/findApiToken [get]
export const findApiToken = (params) => {
  return service({
    url: '/apiToken/findApiToken',
    method: 'get',
    params
  })
}

// @Tags ApiToken
// @Summary 分页获取ApiToken列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ApiToken列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiToken/getApiTokenList [get]
export const getApiTokenList = (params) => {
  return service({
    url: '/apiToken/getApiTokenList',
    method: 'get',
    params
  })
}
