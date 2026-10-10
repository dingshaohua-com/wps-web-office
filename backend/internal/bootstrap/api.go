package bootstrap

import (
	"backend/internal/modules/office"
	"backend/static"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

// RegisterAPI 注册 HTTP 路由及其 OpenAPI 文档，注册过程不会查询数据库；
func RegisterAPI(router *http.ServeMux) {
	router.Handle("/static/", http.StripPrefix("/static/", static.Handler()))

	// 创建一个配置
	config := huma.DefaultConfig("My API", "1.0.0")
	config.CreateHooks = nil
	// 创建humago实例。，并将配置穿进去
	api := humago.New(router, config)
	office.RegisterModule(api)
}
