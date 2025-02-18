package models

import "time"

type KPIDivision struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	ParentID      string    `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
	Parent        KPIParent `bson:"parent,omitempty" json:"parent,omitempty"`
	NameEn        string    `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string    `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string    `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string    `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateKPIDivision struct {
	ParentID      string `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type UpdateKPIDivision struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	ParentID      string `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type ListKPIDivisionResponse struct {
	Count int            `json:"count,omitempty"`
	Items []*KPIDivision `json:"items,omitempty"`
}

type ListKPIDivisionRequest struct {
	ParentID string `json:"parent_id,omitempty"`
	Filter
}

func (k *KPIDivision) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}