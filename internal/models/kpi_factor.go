package models

import "time"

type KPIFactor struct {
	ID            string       `json:"id,omitempty" bson:"_id,omitempty"`
	CriterionID   string       `bson:"criterion_id,omitempty" json:"criterion_id,omitempty"`
	Criterion     KPICriterion `bson:"criterion,omitempty" json:"criterion,omitempty"`
	NameEn        string       `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string       `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	Ratio         float64      `bson:"ratio,omitempty" json:"ratio,omitempty"`
	DescriptionEn string       `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string       `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	CreatedAt     time.Time    `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time    `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateKPIFactor struct {
	CriterionID   string  `bson:"criterion_id,omitempty" json:"criterion_id,omitempty"`
	NameEn        string  `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string  `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	Ratio         float64 `bson:"ratio,omitempty" json:"ratio,omitempty"`
	DescriptionEn string  `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string  `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type UpdateKPIFactor struct {
	ID            string  `json:"id,omitempty" bson:"_id,omitempty"`
	CriterionID   string  `bson:"criterion_id,omitempty" json:"criterion_id,omitempty"`
	NameEn        string  `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string  `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	Ratio         float64 `bson:"ratio,omitempty" json:"ratio,omitempty"`
	DescriptionEn string  `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string  `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type ListKPIFactorResponse struct {
	Count int          `json:"count,omitempty"`
	Items []*KPIFactor `json:"items,omitempty"`
}

type ListKPIFactorRequest struct {
	CriterionID string `json:"criterion_id,omitempty"`
	Filter
}

func (k *KPIFactor) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}