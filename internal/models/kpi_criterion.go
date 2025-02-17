package models

import "time"

type KPICriterion struct {
	ID            string      `json:"id,omitempty" bson:"_id,omitempty"`
	DivisionID    string      `bson:"division_id,omitempty" json:"division_id,omitempty"`
	Division      KPIDivision `bson:"division,omitempty" json:"division,omitempty"`
	NameEn        string      `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string      `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	TotalRatio    float64     `bson:"total_ratio,omitempty" json:"total_ratio,omitempty"`
	DescriptionEn string      `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string      `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	CreatedAt     time.Time   `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time   `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateKPICriterion struct {
	DivisionID    string  `bson:"division_id,omitempty" json:"division_id,omitempty"`
	NameEn        string  `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string  `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	TotalRatio    float64 `bson:"total_ratio,omitempty" json:"total_ratio,omitempty"`
	DescriptionEn string  `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string  `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type UpdateKPICriterion struct {
	ID            string  `json:"id,omitempty" bson:"_id,omitempty"`
	DivisionID    string  `bson:"division_id,omitempty" json:"division_id,omitempty"`
	NameEn        string  `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string  `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	TotalRatio    float64 `bson:"total_ratio,omitempty" json:"total_ratio,omitempty"`
	DescriptionEn string  `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string  `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type ListKPICriterionResponse struct {
	Count int             `json:"count,omitempty"`
	Items []*KPICriterion `json:"items,omitempty"`
}

type ListKPICriterionRequest struct {
	DivisionID string `json:"division_id,omitempty"`
	Filter
}

func (k *KPICriterion) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}