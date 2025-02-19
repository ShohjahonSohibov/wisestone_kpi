package models

import "time"

type KPIFactorIndicator struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	FactorID      string    `bson:"factor_id,omitempty" json:"factor_id,omitempty"`
	Factor        KPIFactor `bson:"factor,omitempty" json:"factor,omitempty"`
	NameEn        string    `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string    `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	ProgressRange int32     `bson:"progress_range,omitempty" json:"progress_range,omitempty"`
	DescriptionEn string    `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string    `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateKPIFactorIndicator struct {
	FactorID      string `bson:"factor_id,omitempty" json:"factor_id,omitempty"`
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	ProgressRange int32  `bson:"progress_range,omitempty" json:"progress_range,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type UpdateKPIFactorIndicator struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	FactorID      string `bson:"factor_id,omitempty" json:"factor_id,omitempty"`
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	ProgressRange int32  `bson:"progress_range,omitempty" json:"progress_range,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type ListKPIFactorIndicatorResponse struct {
	Count int                   `json:"count,omitempty"`
	Items []*KPIFactorIndicator `json:"items,omitempty"`
}

type ListKPIFactorIndicatorRequest struct {
	FactorID string `json:"factor_id,omitempty"`
	Filter
}

func (k *KPIFactorIndicator) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}
