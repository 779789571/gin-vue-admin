package service

import (
	"github.com/779789571/gin-vue-admin/server/service/autocode"
	"github.com/779789571/gin-vue-admin/server/service/example"
	"github.com/779789571/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
