package models

type Filter struct {
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
	Page        int    `json:"page"`
	MultiSearch string `json:"multi_search"`
	SortOrder   string `json:"sort_order" binding:"omitempty,oneof=asc desc"`
}
