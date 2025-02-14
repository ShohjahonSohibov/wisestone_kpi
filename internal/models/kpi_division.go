package models

import "time"

type KPIDivision struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	ParentID      string    `bson:"parent_id" json:"parent_id" binding:"required"`
	Parent        KPIParent `bson:"parent" json:"parent"`
	NameEn        string    `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string    `bson:"name_kr" json:"name_kr" binding:"required"`
	DescriptionEn string    `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string    `bson:"description_kr" json:"description_kr,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateKPIDivision struct {
	ParentID      string `bson:"parent_id" json:"parent_id" binding:"required"`
	NameEn        string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string `bson:"name_kr" json:"name_kr" binding:"required"`
	DescriptionEn string `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr" json:"description_kr,omitempty"`
}

type UpdateKPIDivision struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	ParentID      string `bson:"parent_id" json:"parent_id" binding:"required"`
	NameEn        string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string `bson:"name_kr" json:"name_kr" binding:"required"`
	DescriptionEn string `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr" json:"description_kr,omitempty"`
}

type ListKPIDivisionResponse struct {
	Count int            `json:"count"`
	Items []*KPIDivision `json:"items"`
}

type ListKPIDivisionRequest struct {
	ParentID string `json:"parent_id"`
	Filter
}

func (k *KPIDivision) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}