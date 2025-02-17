package models

import "time"

type KPIFactor struct {
	ID            string       `json:"id,omitempty" bson:"_id,omitempty"`
	CriterionID   string       `bson:"criterion_id" json:"criterion_id"`
	Criterion     KPICriterion `bson:"criterion" json:"criterion"`
	NameEn        string       `bson:"name_en" json:"name_en"`
	NameKr        string       `bson:"name_kr" json:"name_kr"`
	Ratio         float64      `bson:"ratio" json:"ratio"`
	DescriptionEn string       `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string       `bson:"description_kr" json:"description_kr,omitempty"`
	CreatedAt     time.Time    `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateKPIFactor struct {
	CriterionID   string  `bson:"criterion_id" json:"criterion_id"`
	NameEn        string  `bson:"name_en" json:"name_en"`
	NameKr        string  `bson:"name_kr" json:"name_kr"`
	Ratio         float64 `bson:"ratio" json:"ratio"`
	DescriptionEn string  `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string  `bson:"description_kr" json:"description_kr,omitempty"`
}

type UpdateKPIFactor struct {
	ID            string  `json:"id,omitempty" bson:"_id,omitempty"`
	CriterionID   string  `bson:"criterion_id" json:"criterion_id"`
	NameEn        string  `bson:"name_en" json:"name_en"`
	NameKr        string  `bson:"name_kr" json:"name_kr"`
	Ratio         float64 `bson:"ratio" json:"ratio"`
	DescriptionEn string  `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string  `bson:"description_kr" json:"description_kr,omitempty"`
}

type ListKPIFactorResponse struct {
	Count int          `json:"count"`
	Items []*KPIFactor `json:"items"`
}

type ListKPIFactorRequest struct {
	CriterionID string `json:"criterion_id"`
	Filter
}

func (k *KPIFactor) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}