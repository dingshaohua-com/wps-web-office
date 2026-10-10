package api

import (
	"backend/internal/infrastructure/mock"
	"backend/internal/modules/office/application"
	sharedApi "backend/internal/shared/api"
	"context"
	"net/http"
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

func (h *OfficeHandler) FileInfo(_ context.Context, input *FilePathInput) (*sharedApi.Body2[mock.File], error) {
	file, exists, err := mock.GetFile(input.FileID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, sharedApi.NewError(http.StatusNotFound, "文件不存在")
	}

	return sharedApi.NewBody2(file), nil
}

func (h *OfficeHandler) FileDownload(_ context.Context, input *FilePathInput) (*sharedApi.Body2[map[string]string], error) {
	file, exists, err := mock.GetFile(input.FileID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, sharedApi.NewError(http.StatusNotFound, "文件不存在")
	}

	return sharedApi.NewBody2(map[string]string{"url": file.URL}), nil
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

func (h *OfficeHandler) TestFileList(_ context.Context, _ *struct{}) (*sharedApi.Body2[[]mock.File], error) {
	files, err := mock.GetFiles()
	if err != nil {
		return nil, err
	}

	return sharedApi.NewBody2(files), nil
}
