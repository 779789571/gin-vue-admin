package request

import (
	"github.com/779789571/gin-vue-admin/server/model/common/request"
	"github.com/779789571/gin-vue-admin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
