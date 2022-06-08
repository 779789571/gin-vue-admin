import service from '@/utils/request'

// @Tags Rules
// @Summary 创建Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rules true "创建Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rules/createRules [post]
export const createRules = (data) => {
  return service({
    url: '/rules/createRules',
    method: 'post',
    data
  })
}

// @Tags Rules
// @Summary 删除Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rules true "删除Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rules/deleteRules [delete]
export const deleteRules = (data) => {
  return service({
    url: '/rules/deleteRules',
    method: 'delete',
    data
  })
}

// @Tags Rules
// @Summary 删除Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rules/deleteRules [delete]
export const deleteRulesByIds = (data) => {
  return service({
    url: '/rules/deleteRulesByIds',
    method: 'delete',
    data
  })
}

// @Tags Rules
// @Summary 更新Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Rules true "更新Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rules/updateRules [put]
export const updateRules = (data) => {
  return service({
    url: '/rules/updateRules',
    method: 'put',
    data
  })
}

// @Tags Rules
// @Summary 用id查询Rules
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Rules true "用id查询Rules"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rules/findRules [get]
export const findRules = (params) => {
  return service({
    url: '/rules/findRules',
    method: 'get',
    params
  })
}

// @Tags Rules
// @Summary 分页获取Rules列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Rules列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rules/getRulesList [get]
export const getRulesList = (params) => {
  return service({
    url: '/rules/getRulesList',
    method: 'get',
    params
  })
}
