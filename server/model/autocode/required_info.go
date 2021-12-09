// 自动生成模板requiredInfo
package autocode

import (
	"github.com/779789571/gin-vue-admin/server/global"
	"time"
)

// requiredInfo 结构体
// 如果含有time.Time 请自行import time包
type RequiredInfo struct {
	global.GVA_MODEL
	UseRules       string     `json:"useRules" form:"useRules" gorm:"column:use_rules;comment:运行规则;type:varchar(100);"`
	Domain         string     `json:"domain" form:"domain" gorm:"column:domain;comment:;type:varchar(100);"`
	TimeBefore     *time.Time `json:"timeBefore" form:"timeBefore" gorm:"column:time_before;comment:发现时间在..之前;type:date"`
	TimeAfter      *time.Time `json:"timeAfter" form:"timeAfter" gorm:"column:time_after;comment:发现时间在..之后;type:date"`
	TimeAutoCreate *bool      `json:"timeAutoCreate" form:"timeAutoCreate" gorm:"column:time_auto_create;comment:自动生成时间段(当天);type:tinyint"`
	Title          string     `json:"title" form:"title" gorm:"column:title;comment:网站标题;type:varchar(100);"`
	Cert           string     `json:"cert" form:"cert" gorm:"column:cert;comment:证书;type:varchar(100);"`
	Ip             string     `json:"ip" form:"ip" gorm:"column:ip;comment:Ip或ip段;type:varchar(200);"`
	Protocol       string     `json:"protocol" form:"protocol" gorm:"column:protocol;comment:协议;type:varchar(100);"`
	Icp            string     `json:"icp" form:"icp" gorm:"column:icp;comment:icp备案号;type:varchar(100);"`
	WebStatusCode  string     `json:"webStatusCode" form:"webStatusCode" gorm:"column:web_status_code;comment:服务器状态码;type:varchar(100);"`
	Header         string     `json:"header" form:"header" gorm:"column:header;comment:HTTP请求头;type:varchar(100);"`
	Country        string     `json:"country" form:"country" gorm:"column:country;comment:国家;type:varchar(100);"`
	Province       string     `json:"province" form:"province" gorm:"column:province;comment:省份;type:varchar(100);"`
	City           string     `json:"city" form:"city" gorm:"column:city;comment:城市;type:varchar(100);"`
}
