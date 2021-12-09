package autocode

import (
	"github.com/779789571/gin-vue-admin/server/global"
	"github.com/779789571/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/779789571/gin-vue-admin/server/model/autocode/request"
	"github.com/779789571/gin-vue-admin/server/model/common/request"
)

type RulesService struct {
}

// CreateRules 创建Rules记录
// Author [piexlmax](https://github.com/piexlmax)
func (rulesService *RulesService) CreateRules(rules autocode.Rules) (err error) {
	err = global.GVA_DB.Create(&rules).Error
	return err
}

// DeleteRules 删除Rules记录
// Author [piexlmax](https://github.com/piexlmax)
func (rulesService *RulesService) DeleteRules(rules autocode.Rules) (err error) {
	err = global.GVA_DB.Delete(&rules).Error
	return err
}

// DeleteRulesByIds 批量删除Rules记录
// Author [piexlmax](https://github.com/piexlmax)
func (rulesService *RulesService) DeleteRulesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Rules{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRules 更新Rules记录
// Author [piexlmax](https://github.com/piexlmax)
func (rulesService *RulesService) UpdateRules(rules autocode.Rules) (err error) {
	err = global.GVA_DB.Save(&rules).Error
	return err
}

// GetRules 根据id获取Rules记录
// Author [piexlmax](https://github.com/piexlmax)
func (rulesService *RulesService) GetRules(id uint) (err error, rules autocode.Rules) {
	err = global.GVA_DB.Where("id = ?", id).First(&rules).Error
	return
}

// GetRulesInfoList 分页获取Rules记录
// Author [piexlmax](https://github.com/piexlmax)
func (rulesService *RulesService) GetRulesInfoList(info autoCodeReq.RulesSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Rules{})
	var ruless []autocode.Rules
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.RuleName != "" {
		db = db.Where("rule_name LIKE ?", "%"+info.RuleName+"%")
	}
	if info.RuleType != nil {
		db = db.Where("rule_type = ?", info.RuleType)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.IsDynamic != nil {
		db = db.Where("is_dynamic = ?", info.IsDynamic)
	}
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&ruless).Error
	return err, ruless, total
}

func GetValidRulesByType(ruleType int) (err error, list []autocode.Rules) {
	var ruless []autocode.Rules
	db := global.GVA_DB.Model(&autocode.Rules{})
	err = db.Where("rule_type = ?", ruleType).Find(&ruless).Error
	return err, ruless
}
