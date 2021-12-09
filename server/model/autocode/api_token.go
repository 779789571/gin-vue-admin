// 自动生成模板ApiToken
package autocode

import (
	"github.com/779789571/gin-vue-admin/server/global"
)

// ApiToken 结构体
// 如果含有time.Time 请自行import time包
type ApiToken struct {
	global.GVA_MODEL
	ApiType      *int   `json:"apiType" form:"apiType" gorm:"column:api_type;comment:API类型;type:int"`
	Status       *int   `json:"status" form:"status" gorm:"column:status;comment:API状态;type:int"`
	Account      string `json:"account" form:"account" gorm:"column:account;comment:账号(如有);type:varchar(100);"`
	Content      string `json:"content" form:"content" gorm:"column:content;comment:apitoken内容;type:varchar(100);"`
	LimitTimes   *int   `json:"limitTimes" form:"limitTimes" gorm:"column:limit_times;comment:次数限制;type:int"`
	RemaingTimes *int   `json:"remaingTimes" form:"remaingTimes" gorm:"column:remaing_times;comment:剩余次数;type:int"`
	Desc         string `json:"desc" form:"desc" gorm:"column:desc;comment:备注;type:varchar(300);"`
}
