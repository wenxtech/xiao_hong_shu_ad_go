package models

// CommonResponse struct for CommonResponse
type CommonResponse struct {
	Code      *int64                 `json:"code,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Message   *string                `json:"message,omitempty"`
	RequestId *string                `json:"request_id,omitempty"`
}
