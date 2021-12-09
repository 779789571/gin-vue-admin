package request

import (
	"github.com/779789571/gin-vue-admin/server/model/autocode"
	"github.com/779789571/gin-vue-admin/server/model/common/request"
)

type RulesSearch struct {
	autocode.Rules
	request.PageInfo
}
