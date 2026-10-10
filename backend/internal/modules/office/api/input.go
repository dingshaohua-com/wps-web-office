package api

import (
	"backend/internal/shared/network"

	"github.com/danielgtaylor/huma/v2"
)

type FilePathInput struct {
	FileID string `path:"file_id"`
}

type FileDownloadInput struct {
	FileID        string `path:"file_id"`
	RequestOrigin network.RequestOrigin
}

// Resolve 是 Huma 的请求参数解析钩子，方法名由 huma.Resolver 接口规定。
func (input *FileDownloadInput) Resolve(ctx huma.Context) []error {
	input.RequestOrigin = network.NewRequestOrigin(ctx)
	return nil
}
