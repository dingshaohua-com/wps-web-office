package office

import (
	"backend/internal/modules/office/api"
	"backend/internal/modules/office/application"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterModule(serverApi huma.API) {
	// 读路径依赖组装 (Query Side - 直连 db，没有任何其他依赖！)
	querySvc := application.NewQueryOfficeService() // 推荐把结构体/构造命名为 Service 或 QueryService

	// API 层同时注入【写服务】与【读服务】
	handler := api.NewOfficeHandler(querySvc)

	api.RegisterRoutes(handler, serverApi)
}
