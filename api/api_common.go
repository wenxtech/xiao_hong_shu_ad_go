package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	. "github.com/wenxtech/xiao_hong_shu_ad_go/models"
)

type CommonApiService service

type CommonApiRequest struct {
	ctx          context.Context
	path         string
	ApiService   *CommonApiService
	requestBody  interface{}
	requestQuery map[string]interface{}
	requestFile  []FormFile
	requestForm  map[string]interface{}
	httpMethod   string
	contentType  string
}

func (r *CommonApiRequest) WithLog(enable bool) *CommonApiRequest {
	if enable {
		r.ctx = context.WithValue(r.ctx, ContextEnableLog, true)
	}
	return r
}

func (r *CommonApiRequest) AccessToken(accessToken string) *CommonApiRequest {
	r.ctx = context.WithValue(r.ctx, ContextAccessToken, accessToken)
	return r
}

func (r *CommonApiRequest) RequestBody(requestBody interface{}) *CommonApiRequest {
	r.requestBody = requestBody
	return r
}

func (r *CommonApiRequest) RequestQuery(requestQuery map[string]interface{}) *CommonApiRequest {
	r.requestQuery = requestQuery
	return r
}

func (r *CommonApiRequest) RequestFile(requestFile []FormFile) *CommonApiRequest {
	r.requestFile = requestFile
	return r
}

func (r *CommonApiRequest) RequestForm(requestForm map[string]interface{}) *CommonApiRequest {
	r.requestForm = requestForm
	return r
}

func (r *CommonApiRequest) Execute() (*CommonResponse, *http.Response, error) {
	return r.ApiService.execute(r)
}

func (a *CommonApiService) Get(ctx context.Context, path string) *CommonApiRequest {
	return &CommonApiRequest{
		ctx:        ctx,
		path:       path,
		ApiService: a,
		httpMethod: http.MethodGet,
	}
}

func (a *CommonApiService) Post(ctx context.Context, path string) *CommonApiRequest {
	return &CommonApiRequest{
		ctx:         ctx,
		path:        path,
		ApiService:  a,
		httpMethod:  http.MethodPost,
		contentType: "application/json",
	}
}

func (a *CommonApiService) PostMultipart(ctx context.Context, path string) *CommonApiRequest {
	return &CommonApiRequest{
		ctx:         ctx,
		path:        path,
		ApiService:  a,
		httpMethod:  http.MethodPost,
		contentType: "multipart/form-data",
	}
}

func (a *CommonApiService) execute(r *CommonApiRequest) (*CommonResponse, *http.Response, error) {
	var localVarReturnValue *CommonResponse

	r.ctx = a.client.prepareCtx(r.ctx)
	localBasePath := a.client.Cfg.GetBasePath()
	localVarPath := localBasePath + r.path
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.contentType != "" {
		localVarHeaderParams["Content-Type"] = r.contentType
	}
	if len(r.requestQuery) > 0 {
		for k, v := range r.requestQuery {
			parameterAddToHeaderOrQuery(localVarQueryParams, k, v, "")
		}
	}
	if len(r.requestForm) > 0 {
		for k, v := range r.requestForm {
			parameterAddToHeaderOrQuery(localVarFormParams, k, v, "")
		}
	}

	req, err := a.client.prepareRequest(r.ctx, localVarPath, r.httpMethod, r.requestBody,
		localVarHeaderParams, localVarQueryParams, localVarFormParams, r.requestFile)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r.ctx, req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
