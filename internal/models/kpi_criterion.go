package models

import "time"

type KPICriterion struct {
	ID            string      `json:"id,omitempty" bson:"_id,omitempty"`
	DivisionID    string      `bson:"division_id" json:"division_id" binding:"required"`
	Division      KPIDivision `bson:"division" json:"division"`
	NameEn        string      `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string      `bson:"name_kr" json:"name_kr" binding:"required"`
	TotalRatio    float64     `bson:"total_ratio" json:"total_ratio" binding:"required"`
	DescriptionEn string      `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string      `bson:"description_kr" json:"description_kr,omitempty"`
	CreatedAt     time.Time   `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateKPICriterion struct {
	DivisionID    string  `bson:"division_id" json:"division_id" binding:"required"`
	NameEn        string  `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string  `bson:"name_kr" json:"name_kr" binding:"required"`
	TotalRatio    float64 `bson:"total_ratio" json:"total_ratio" binding:"required"`
	DescriptionEn string  `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string  `bson:"description_kr" json:"description_kr,omitempty"`
}

type UpdateKPICriterion struct {
	ID            string  `json:"id,omitempty" bson:"_id,omitempty"`
	DivisionID    string  `bson:"division_id" json:"division_id" binding:"required"`
	NameEn        string  `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string  `bson:"name_kr" json:"name_kr" binding:"required"`
	TotalRatio    float64 `bson:"total_ratio" json:"total_ratio" binding:"required"`
	DescriptionEn string  `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string  `bson:"description_kr" json:"description_kr,omitempty"`
}

type ListKPICriterionResponse struct {
	Count int             `json:"count"`
	Items []*KPICriterion `json:"items"`
}

type ListKPICriterionRequest struct {
	DivisionID string `json:"division_id"`
	Filter
}

func (k *KPICriterion) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}