package autocode

import (
	"github.com/779789571/gin-vue-admin/server/global"
	"github.com/779789571/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/779789571/gin-vue-admin/server/model/autocode/request"
	"github.com/779789571/gin-vue-admin/server/model/common/request"
)

type CollectorService struct {
}

// CreateCollector 创建Collector记录
// Author [piexlmax](https://github.com/piexlmax)
func (collectorService *CollectorService) CreateCollector(collector autocode.Collector) (err error) {
	err = global.GVA_DB.Create(&collector).Error
	return err
}

// DeleteCollector 删除Collector记录
// Author [piexlmax](https://github.com/piexlmax)
func (collectorService *CollectorService) DeleteCollector(collector autocode.Collector) (err error) {
	err = global.GVA_DB.Delete(&collector).Error
	return err
}

// DeleteCollectorByIds 批量删除Collector记录
// Author [piexlmax](https://github.com/piexlmax)
func (collectorService *CollectorService) DeleteCollectorByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Collector{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCollector 更新Collector记录
// Author [piexlmax](https://github.com/piexlmax)
func (collectorService *CollectorService) UpdateCollector(collector autocode.Collector) (err error) {
	err = global.GVA_DB.Save(&collector).Error
	return err
}

// GetCollector 根据id获取Collector记录
// Author [piexlmax](https://github.com/piexlmax)
func (collectorService *CollectorService) GetCollector(id uint) (err error, collector autocode.Collector) {
	err = global.GVA_DB.Where("id = ?", id).First(&collector).Error
	return
}

// GetCollectorInfoList 分页获取Collector记录
// Author [piexlmax](https://github.com/piexlmax)
func (collectorService *CollectorService) GetCollectorInfoList(info autoCodeReq.CollectorSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Collector{})
	var collectors []autocode.Collector
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Domain != "" {
		db = db.Where("domain LIKE ?", "%"+info.Domain+"%")
	}
	if info.BelongToSrc != "" {
		db = db.Where("belong_to_src = ?", info.BelongToSrc)
	}
	if info.IsWildDomain != nil {
		db = db.Where("is_wild_domain = ?", info.IsWildDomain)
	}
	if info.DiscoveryTime != nil {
		db = db.Where("discovery_time > ?", info.DiscoveryTime)
	}
	if info.IpAddress != "" {
		db = db.Where("ip_address LIKE ?", "%"+info.IpAddress+"%")
	}
	if info.OpenPorts != "" {
		db = db.Where("open_ports LIKE ?", "%"+info.OpenPorts+"%")
	}
	if info.IsMonitoring != nil {
		db = db.Where("is_monitoring = ?", info.IsMonitoring)
	}
	if info.WebStatusCode != nil {
		db = db.Where("web_status_code = ?", info.WebStatusCode)
	}
	if info.WebTitles != "" {
		db = db.Where("web_titles LIKE ?", "%"+info.WebTitles+"%")
	}
	if info.HadScanned != nil {
		db = db.Where("had_scanned = ?", info.HadScanned)
	}
	if info.Middleware != "" {
		db = db.Where("middleware LIKE ?", "%"+info.Middleware+"%")
	}
	if info.Fingerprint != "" {
		db = db.Where("fingerprint LIKE ?", "%"+info.Fingerprint+"%")
	}
	if info.WebPageSize != "" {
		db = db.Where("web_page_size <> ?", info.WebPageSize)
	}
	if info.PrimaryDomain != "" {
		db = db.Where("primary_domain = ?", info.PrimaryDomain)
	}
	if info.IsOnBlacklist != nil {
		db = db.Where("is_on_blacklist = ?", info.IsOnBlacklist)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&collectors).Error
	return err, collectors, total
}
