/*
小红书AD

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// JgDataReportOfflineAccountPostRequest struct for JgDataReportOfflineAccountPostRequest
type JgDataReportOfflineAccountPostRequest struct {
	AdvertiserId int64 `json:"advertiser_id"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	TimeUnit *string `json:"time_unit,omitempty"`
	MarketingTarget []int64 `json:"marketing_target,omitempty"`
	BiddingStrategy []int64 `json:"bidding_strategy,omitempty"`
	OptimizeTarget []int64 `json:"optimize_target,omitempty"`
	Placement []int64 `json:"placement,omitempty"`
	PromotionTarget []int64 `json:"promotion_target,omitempty"`
	DeliveryMode []int64 `json:"delivery_mode,omitempty"`
	SplitColumns []string `json:"split_columns,omitempty"`
	SortColumn *string `json:"sort_column,omitempty"`
	Sort *string `json:"sort,omitempty"`
	PageNum *int64 `json:"page_num,omitempty"`
	PageSize *int64 `json:"page_size,omitempty"`
	DataCaliber *int64 `json:"data_caliber,omitempty"`
}

type _JgDataReportOfflineAccountPostRequest JgDataReportOfflineAccountPostRequest


