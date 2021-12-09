// 自动生成模板Rules
package autocode

import (
	"github.com/779789571/gin-vue-admin/server/global"
)

// Rules 结构体
// 如果含有time.Time 请自行import time包
type Rules struct {
	global.GVA_MODEL
	RuleName  string `json:"ruleName" form:"ruleName" gorm:"column:rule_name;comment:规则名称;type:varchar(100);"`
	RuleType  *int   `json:"ruleType" form:"ruleType" gorm:"column:rule_type;comment:rule类型;type:int"`
	Status    *int   `json:"status" form:"status" gorm:"column:status;comment:rule状态;type:int"`
	IsDynamic *bool  `json:"isDynamic" form:"isDynamic" gorm:"column:is_dynamic;comment:是否为动态规则;type:tinyint"`
	Content   string `json:"content" form:"content" gorm:"column:content;comment:rule内容;type:varchar(300);"`
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:备注;type:varchar(300);"`
}

// TableName Rules 表名
func (Rules) TableName() string {
	return "rules"
}
