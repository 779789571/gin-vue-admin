package response

import "github.com/779789571/gin-vue-admin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
