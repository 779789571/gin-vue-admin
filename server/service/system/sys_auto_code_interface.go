package system

import (
	"github.com/779789571/gin-vue-admin/server/global"
	"github.com/779789571/gin-vue-admin/server/model/system/response"
)

type Database interface {
	GetDB() (data []response.Db, err error)
	GetTables(dbName string) (data []response.Table, err error)
	GetColumn(tableName string, dbName string) (data []response.Column, err error)
}

func (autoCodeService *AutoCodeService) Database() Database {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return AutoCodeMysql
	case "pgsql":
		return AutoCodePgsql
	default:
		return AutoCodeMysql
	}
}
