package middleware

import (
	"context"
	"net/http"

	"github.com/wenxtech/xiao_hong_shu_ad_go/api"
)

func AuthMiddleware(next api.Endpoint) api.Endpoint {
	return func(ctx context.Context, req *http.Request) (resp *http.Response, err error) {
		token := ctx.Value(api.ContextAccessToken)
		if token != nil {
			if tokenStr, ok := token.(string); ok {
				req.Header.Set("Access-Token", tokenStr)
			}
		}
		return next(ctx, req)
	}
}
