package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ApiTokenService struct {
}

// CreateApiToken 创建ApiToken记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiTokenService *ApiTokenService) CreateApiToken(apiToken autocode.ApiToken) (err error) {
	err = global.GVA_DB.Create(&apiToken).Error
	return err
}

// DeleteApiToken 删除ApiToken记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiTokenService *ApiTokenService) DeleteApiToken(apiToken autocode.ApiToken) (err error) {
	err = global.GVA_DB.Delete(&apiToken).Error
	return err
}

// DeleteApiTokenByIds 批量删除ApiToken记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiTokenService *ApiTokenService) DeleteApiTokenByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ApiToken{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateApiToken 更新ApiToken记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiTokenService *ApiTokenService) UpdateApiToken(apiToken autocode.ApiToken) (err error) {
	err = global.GVA_DB.Save(&apiToken).Error
	return err
}

// GetApiToken 根据id获取ApiToken记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiTokenService *ApiTokenService) GetApiToken(id uint) (err error, apiToken autocode.ApiToken) {
	err = global.GVA_DB.Where("id = ?", id).First(&apiToken).Error
	return
}

// GetApiTokenInfoList 分页获取ApiToken记录
// Author [piexlmax](https://github.com/piexlmax)
func (apiTokenService *ApiTokenService) GetApiTokenInfoList(info autoCodeReq.ApiTokenSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.ApiToken{})
	var apiTokens []autocode.ApiToken
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ApiType != nil {
		db = db.Where("api_type = ?", info.ApiType)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.LimitTimes != nil {
		db = db.Where("limit_times > ?", info.LimitTimes)
	}
	if info.RemaingTimes != nil {
		db = db.Where("remaing_times > ?", info.RemaingTimes)
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&apiTokens).Error
	return err, apiTokens, total
}
