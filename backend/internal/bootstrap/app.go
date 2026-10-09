package bootstrap

import (
	sharedApi "backend/internal/shared/api"
	"backend/internal/shared/network"
	"backend/internal/webui"
	"log"
	"net/http"
	"strings"

	"backend/internal/infrastructure"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Router   *http.ServeMux
	Config   *infrastructure.Config
	Database *gorm.DB
	Server   *http.Server
	Redis    *redis.Client
}

func (app *App) Run() {
	addr := ":" + app.Config.HTTPPort
	server := &http.Server{
		Addr:    addr,
		Handler: sharedApi.Cors(NewHTTPHandler(app.Router, webui.Handler())),
	}
	var serviceURLs, docURLs []string
	for _, host := range network.LocalLANIPv4Addresses() {
		serviceURLs = append(serviceURLs, "http://"+host+addr)
		docURLs = append(docURLs, "http://"+host+addr+"/docs")
	}
	log.Printf("HTTP已服务启动: %s", strings.Join(serviceURLs, ", "))
	log.Printf("OpenAPI文档: %s", strings.Join(docURLs, ", "))
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("HTTP 服务启动失败: %v", err)
	}
}

// NewApp 这里负责组装整个系统
func NewApp() (*App, error) {
	// 加载环境变量
	cfg := infrastructure.LoadConfig()

	// 初始化服务
	router := http.NewServeMux() // http.NewServeMux() 用来创建一个 HTTP 路由器，也就是根据请求路径，把请求分发给对应的处理函数
	RegisterAPI(router)

	return &App{
		Router: router,
		Config: cfg,
	}, nil
}
