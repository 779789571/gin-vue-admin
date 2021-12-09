package response

import (
	"github.com/779789571/gin-vue-admin/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
