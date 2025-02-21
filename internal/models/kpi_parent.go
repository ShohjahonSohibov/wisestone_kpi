package models

import (
	"time"
)

type KPIParent struct {
	ID             string             `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn         string             `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr         string             `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn  string             `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr  string             `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	Year           string             `bson:"year,omitempty" json:"year,omitempty"`
	Status         string             `bson:"status,omitempty" json:"status,omitempty"`
	Type           string             `bson:"type,omitempty" json:"type,omitempty"`
	RejectionCount int                `bson:"rejection_count,omitempty" json:"rejection_count,omitempty"`
	CreatedAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Divisions      []ShortKPIDivision `bson:"divisions,omitempty" json:"divisions,omitempty"`
}

type CreateKPIParent struct {
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	Type          string `bson:"type,omitempty" json:"type,omitempty"`
}

type UpdateKPIParent struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	Type          string `bson:"type,omitempty" json:"type,omitempty"`
}

type UpdateKPIParentStatus struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Status string `bson:"status,omitempty" json:"status,omitempty"`
}

type ListKPIParentResponse struct {
	Count int          `json:"count,omitempty"`
	Items []*KPIParent `json:"items,omitempty"`
}

type ListKPIParentRequest struct {
	Filter
	Year   string `json:"year"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

type ShortKPIDivision struct {
	ID            string              `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string              `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string              `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string              `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string              `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	Criterions    []ShortKPICriterion `bson:"criterions,omitempty" json:"criterions,omitempty"`
}

type ShortKPICriterion struct {
	ID            string            `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string            `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string            `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	TotalRatio    float64           `bson:"total_ratio,omitempty" json:"total_ratio,omitempty"`
	DescriptionEn string            `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string            `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	Factors       []*ShortKPIFactor `bson:"factors,omitempty" json:"factors,omitempty"`
}

type ShortKPIFactor struct {
	ID               string                     `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn           string                     `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr           string                     `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	Ratio            float64                    `bson:"ratio,omitempty" json:"ratio,omitempty"`
	DescriptionEn    string                     `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr    string                     `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	FactorIndicators []*ShortKPIFactorIndicator `bson:"factor_indicators,omitempty" json:"factor_indicators,omitempty"`
}

type ShortKPIFactorIndicator struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	ProgressRange int32  `bson:"progress_range,omitempty" json:"progress_range,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

func (k *KPIParent) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}
