package autocode

import (
	"github.com/779789571/gin-vue-admin/server/global"
	"github.com/779789571/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/779789571/gin-vue-admin/server/model/autocode/request"
	"github.com/779789571/gin-vue-admin/server/model/common/request"
)

type RequiredInfoService struct {
}

// CreaterequiredInfo 创建requiredInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (requiredInfoService *RequiredInfoService) CreaterequiredInfo(requiredInfo autocode.RequiredInfo) (err error) {
	err = global.GVA_DB.Create(&requiredInfo).Error
	return err
}

// DeleterequiredInfo 删除requiredInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (requiredInfoService *RequiredInfoService) DeleterequiredInfo(requiredInfo autocode.RequiredInfo) (err error) {
	err = global.GVA_DB.Delete(&requiredInfo).Error
	return err
}

// DeleterequiredInfoByIds 批量删除requiredInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (requiredInfoService *RequiredInfoService) DeleterequiredInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.RequiredInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdaterequiredInfo 更新requiredInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (requiredInfoService *RequiredInfoService) UpdaterequiredInfo(requiredInfo autocode.RequiredInfo) (err error) {
	err = global.GVA_DB.Save(&requiredInfo).Error
	return err
}

// GetrequiredInfo 根据id获取requiredInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (requiredInfoService *RequiredInfoService) GetrequiredInfo(id uint) (err error, requiredInfo autocode.RequiredInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&requiredInfo).Error
	return
}

// GetrequiredInfoInfoList 分页获取requiredInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (requiredInfoService *RequiredInfoService) GetrequiredInfoInfoList(info autoCodeReq.RequiredInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.RequiredInfo{})
	var requiredInfos []autocode.RequiredInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UseRules != "" {
		db = db.Where("use_rules LIKE ?", "%"+info.UseRules+"%")
	}
	if info.Domain != "" {
		db = db.Where("domain LIKE ?", "%"+info.Domain+"%")
	}
	if info.TimeBefore != nil {
		db = db.Where("time_before < ?", info.TimeBefore)
	}
	if info.TimeAfter != nil {
		db = db.Where("time_after > ?", info.TimeAfter)
	}
	if info.TimeAutoCreate != nil {
		db = db.Where("time_auto_create = ?", info.TimeAutoCreate)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Cert != "" {
		db = db.Where("cert LIKE ?", "%"+info.Cert+"%")
	}
	if info.Ip != "" {
		db = db.Where("ip LIKE ?", "%"+info.Ip+"%")
	}
	if info.Protocol != "" {
		db = db.Where("protocol LIKE ?", "%"+info.Protocol+"%")
	}
	if info.Icp != "" {
		db = db.Where("icp LIKE ?", "%"+info.Icp+"%")
	}
	if info.WebStatusCode != "" {
		db = db.Where("web_status_code LIKE ?", "%"+info.WebStatusCode+"%")
	}
	if info.Header != "" {
		db = db.Where("header LIKE ?", "%"+info.Header+"%")
	}
	if info.Country != "" {
		db = db.Where("country LIKE ?", "%"+info.Country+"%")
	}
	if info.Province != "" {
		db = db.Where("province LIKE ?", "%"+info.Province+"%")
	}
	if info.City != "" {
		db = db.Where("city LIKE ?", "%"+info.City+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&requiredInfos).Error
	return err, requiredInfos, total
}

func GetRulesByRuleName(ruleName string) (err error, requiredInfos []autocode.RequiredInfo) {
	//var requiredInfos []autocode.RequiredInfo
	db := global.GVA_DB.Model(&autocode.RequiredInfo{})
	db = db.Where("use_rules LIKE ?", "%"+ruleName+"%")
	err = db.Find(&requiredInfos).Error
	return err, requiredInfos
}
