package middleware

import (
	"context"
	"net/http"
	"runtime"

	"github.com/wenxtech/xiao_hong_shu_ad_go/api"
)

func HeaderMiddleware(next api.Endpoint) api.Endpoint {
	return func(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
		req.Header.Set("X-Sdk-Language", "go")
		req.Header.Set("X-Sdk-Language-Version", runtime.Version())
		req.Header.Set("X-Sdk-Version", api.Version)
		return next(ctx, req)
	}
}
