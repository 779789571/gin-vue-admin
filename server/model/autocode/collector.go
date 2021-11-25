// 自动生成模板Collector
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Collector 结构体
// 如果含有time.Time 请自行import time包
type Collector struct {
	global.GVA_MODEL
	Domain        string     `json:"domain" form:"domain" gorm:"column:domain;comment:域名;type:varchar(200);"`
	BelongToSrc   string     `json:"belongToSrc" form:"belongToSrc" gorm:"column:belong_to_src;comment:src归属;type:varchar(100);"`
	IsWildDomain  *bool      `json:"isWildDomain" form:"isWildDomain" gorm:"column:is_wild_domain;comment:是否为泛解析;type:tinyint"`
	DiscoveryTime *time.Time `json:"discoveryTime" form:"discoveryTime" gorm:"column:discovery_time;comment:域名发现时间;type:datetime"`
	IpAddress     string     `json:"ipAddress" form:"ipAddress" gorm:"column:ip_address;comment:IP地址;type:varchar(100);"`
	OpenPorts     string     `json:"openPorts" form:"openPorts" gorm:"column:open_ports;comment:开放端口;type:varchar(999);"`
	IsMonitoring  *bool      `json:"isMonitoring" form:"isMonitoring" gorm:"column:is_monitoring;comment:是否监控该域名;type:tinyint"`
	WebStatusCode *int       `json:"webStatusCode" form:"webStatusCode" gorm:"column:web_status_code;comment:web响应码;type:int"`
	WebTitles     string     `json:"webTitles" form:"webTitles" gorm:"column:web_titles;comment:web页面标题;type:varchar(100);"`
	HadScanned    *bool      `json:"hadScanned" form:"hadScanned" gorm:"column:had_scanned;comment:是否已扫描;type:tinyint"`
	Middleware    string     `json:"middleware" form:"middleware" gorm:"column:middleware;comment:中间件信息;type:varchar(999);"`
	Fingerprint   string     `json:"fingerprint" form:"fingerprint" gorm:"column:fingerprint;comment:指纹信息;type:varchar(999);"`
	WebPageSize   string     `json:"webPageSize" form:"webPageSize" gorm:"column:web_page_size;comment:页面大小;type:varchar(100);"`
	PrimaryDomain string     `json:"primaryDomain" form:"primaryDomain" gorm:"column:primary_domain;comment:一级域名;type:varchar(100);"`
	IsOnBlacklist *bool      `json:"isOnBlacklist" form:"isOnBlacklist" gorm:"column:is_on_blacklist;comment:是否在黑名单中;type:tinyint"`
}

// TableName Collector 表名
func (Collector) TableName() string {
	return "collector"
}
