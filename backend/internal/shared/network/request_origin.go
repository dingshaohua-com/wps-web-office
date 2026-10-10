package network

import (
	"net/url"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

// RequestOrigin 从当前请求中解析协议和主机，用于生成绝对 URL。
type RequestOrigin struct {
	baseURL string
}

func NewRequestOrigin(ctx huma.Context) RequestOrigin {
	scheme := firstForwardedValue(ctx.Header("X-Forwarded-Proto"))
	if scheme != "http" && scheme != "https" {
		scheme = "http"
		if ctx.TLS() != nil {
			scheme = "https"
		}
	}

	host := firstForwardedValue(ctx.Header("X-Forwarded-Host"))
	if host == "" {
		host = ctx.Host()
	}

	return RequestOrigin{baseURL: scheme + "://" + host}
}

// BuildAbsoluteURL 将相对路径转换为基于当前请求 Host 的绝对 URL。
func (origin RequestOrigin) BuildAbsoluteURL(reference string) string {
	base, baseErr := url.Parse(origin.baseURL)
	target, targetErr := url.Parse(reference)
	if baseErr != nil || targetErr != nil {
		return reference
	}
	return base.ResolveReference(target).String()
}

func firstForwardedValue(value string) string {
	value, _, _ = strings.Cut(value, ",")
	return strings.TrimSpace(value)
}
