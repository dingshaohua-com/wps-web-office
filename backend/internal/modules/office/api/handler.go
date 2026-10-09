package api

import (
	"backend/internal/modules/office/application"
	sharedApi "backend/internal/shared/api"
	"context"
)

type FilePathInput struct {
	FileID string `path:"file_id"`
}

type OfficeHandler struct {
	service *application.QueryOfficeService
}

func NewOfficeHandler(service *application.QueryOfficeService) *OfficeHandler {
	return &OfficeHandler{service: service}
}

func (h *OfficeHandler) FileInfo(ctx context.Context, input *FilePathInput) (*sharedApi.Body[map[string]any], error) {
	info := map[string]any{
		"code": 0,
		"data": map[string]any{
			"id":          "jl_docx",
			"name":        "jl.docx",
			"version":     1,
			"size":        118221,
			"create_time": 1791532236,
			"modify_time": 1791532236,
			"creator_id":  "system",
			"modifier_id": "system",
		},
	}
	return sharedApi.NewBody(info), nil
}

func (h *OfficeHandler) FileDownload(ctx context.Context, input *FilePathInput) (*sharedApi.Body[map[string]any], error) {

	info := map[string]any{
		"code": 0,
		"data": map[string]any{
			"url": "https://dsh-1300009960.cos.ap-beijing.myqcloud.com/office/jl.docx",
		},
	}
	return sharedApi.NewBody(info), nil
}

func (h *OfficeHandler) FilePermission(ctx context.Context, input *FilePathInput) (*sharedApi.Body[map[string]any], error) {

	info := map[string]any{
		"code": 0,
		"data": map[string]any{
			"read":     1,
			"update":   0,
			"download": 1,
			"copy":     1,
			"print":    1,
			"rename":   0,
			"history":  0,
			"saveas":   0,
			"comment":  0,
		},
	}
	return sharedApi.NewBody(info), nil
}
