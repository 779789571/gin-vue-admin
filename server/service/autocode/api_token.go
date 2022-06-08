package autocode

import (
	"errors"
	"fmt"
	"github.com/779789571/gin-vue-admin/server/global"
	"github.com/779789571/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/779789571/gin-vue-admin/server/model/autocode/request"
	"github.com/779789571/gin-vue-admin/server/model/common/request"
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
	if info.Account != "" {
		db = db.Where("account LIKE ?", "%"+info.Account+"%")
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

//map[string]string map[token]备注 备注相同为同组key与secret
func GetApiToken(api_type int) (error, interface{}, interface{}, interface{}) {
	var apiTokens []autocode.ApiToken
	db := global.GVA_DB.Model(&autocode.ApiToken{})
	db = db.Where("api_type = ?", api_type).Where("status = ?", 1)
	err := db.Find(&apiTokens).Error
	if err != nil {
		return err, nil, nil, nil
	}
	if len(apiTokens) == 0 {
		return errors.New(fmt.Sprintf("没找到%v 的token信息", api_type)), "", "", 0
	}
	if len(apiTokens) == 1 {
		return nil, apiTokens[0].Account, apiTokens[0].Content, *apiTokens[0].LimitTimes
	}
	var accountList []string
	var keyList []string
	var limitTimesList []int
	for _, v := range apiTokens {
		if v.Account != "" {
			accountList = append(accountList, v.Account)
		}
		keyList = append(keyList, v.Content)
		limitTimesList = append(limitTimesList, *v.LimitTimes)
	}
	return nil, accountList, keyList, limitTimesList
}
