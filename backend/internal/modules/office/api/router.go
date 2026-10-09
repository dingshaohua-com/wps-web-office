package api

import (
	sharedApi "backend/internal/shared/api"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(handler *OfficeHandler, api huma.API) {
	feedGroup := sharedApi.NewGroup(api, "/office/v3/3rd/files", "office")
	feedGroup.UseSimpleModifier(func(op *huma.Operation) {
		op.Description = "文档服务"
	})

	huma.Get(feedGroup, "/{file_id}", handler.FileInfo, func(op *huma.Operation) {
		op.OperationID = "file_info"
		op.Summary = "文件信息"
	})

	huma.Get(feedGroup, "/{file_id}/download", handler.FileDownload, func(op *huma.Operation) {
		op.OperationID = "file_download"
		op.Summary = "文件下载"
	})

	huma.Get(feedGroup, "/{file_id}/permission", handler.FilePermission, func(op *huma.Operation) {
		op.OperationID = "file_permission"
		op.Summary = "文件权限"
	})
}
