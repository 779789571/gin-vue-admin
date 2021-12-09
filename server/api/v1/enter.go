package v1

import (
	"github.com/779789571/gin-vue-admin/server/api/v1/autocode"
	"github.com/779789571/gin-vue-admin/server/api/v1/example"
	"github.com/779789571/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
