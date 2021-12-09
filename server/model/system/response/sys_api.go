package response

import "github.com/779789571/gin-vue-admin/server/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}
