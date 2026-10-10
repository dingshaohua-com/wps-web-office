package api

import (
	sharedApi "backend/internal/shared/api"

	"github.com/danielgtaylor/huma/v2"
)

func operation(id, summary string) func(*huma.Operation) {
	return func(op *huma.Operation) {
		op.OperationID = id
		op.Summary = summary
	}
}

func RegisterRoutes(handler *OfficeHandler, api huma.API) {
	officeGroup := sharedApi.NewGroup(api, "/office/v3/3rd/files", "office")
	huma.Get(officeGroup, "/{file_id}", handler.FileInfo, operation("file_info", "文件信息"))
	huma.Get(officeGroup, "/{file_id}/download", handler.FileDownload, operation("file_download", "文件下载"))
	huma.Get(officeGroup, "/{file_id}/permission", handler.FilePermission, operation("file_permission", "文件权限"))

	testOfficeGroup := sharedApi.NewGroup(api, "/test-office", "test-office")
	huma.Get(testOfficeGroup, "", handler.TestFileList, operation("test_file_list", "测试文件列表"))
}
