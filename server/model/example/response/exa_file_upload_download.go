package response

import "github.com/779789571/gin-vue-admin/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
