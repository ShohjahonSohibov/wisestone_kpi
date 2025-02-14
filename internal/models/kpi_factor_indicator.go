package models

import "time"

type KPIFactorIndicator struct {
	ID             string    `json:"id,omitempty" bson:"_id,omitempty"`
	FactorID       string    `bson:"factor_id" json:"factor_id" binding:"required"`
	Factor         KPIFactor `bson:"factor" json:"factor"`
	NameEn         string    `bson:"name_en" json:"name_en" binding:"required"`
	NameKr         string    `bson:"name_kr" json:"name_kr" binding:"required"`
	ProgressRange  string    `bson:"progress_range" json:"progress_range" binding:"required"`
	DescriptionEn  string    `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr  string    `bson:"description_kr" json:"description_kr,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateKPIFactorIndicator struct {
	FactorID      string `bson:"factor_id" json:"factor_id" binding:"required"`
	NameEn        string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string `bson:"name_kr" json:"name_kr" binding:"required"`
	ProgressRange string `bson:"progress_range" json:"progress_range" binding:"required"`
	DescriptionEn string `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr" json:"description_kr,omitempty"`
}

type UpdateKPIFactorIndicator struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	FactorID      string `bson:"factor_id" json:"factor_id" binding:"required"`
	NameEn        string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string `bson:"name_kr" json:"name_kr" binding:"required"`
	ProgressRange string `bson:"progress_range" json:"progress_range" binding:"required"`
	DescriptionEn string `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr" json:"description_kr,omitempty"`
}

type ListKPIFactorIndicatorResponse struct {
	Count int                   `json:"count"`
	Items []*KPIFactorIndicator `json:"items"`
}

type ListKPIFactorIndicatorRequest struct {
	FactorID string `json:"factor_id"`
	Filter
}

func (k *KPIFactorIndicator) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}