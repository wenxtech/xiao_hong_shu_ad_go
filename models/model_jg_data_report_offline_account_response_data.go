/*
小红书AD

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// JgDataReportOfflineAccountResponseData struct for JgDataReportOfflineAccountResponseData
type JgDataReportOfflineAccountResponseData struct {
	TotalCount int64 `json:"total_count"`
	AggregationData JgDataReportOfflineAccountResponseDataAggregationData `json:"aggregation_data"`
	DataList []JgDataReportOfflineAccountResponseDataDataListInner `json:"data_list"`
}

type _JgDataReportOfflineAccountResponseData JgDataReportOfflineAccountResponseData


