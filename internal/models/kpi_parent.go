package models

import "time"

type KPIParent struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string    `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string    `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string    `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string    `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
	Year          string    `bson:"year,omitempty" json:"year,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CreateKPIParent struct {
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type UpdateKPIParent struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string `bson:"name_en,omitempty" json:"name_en,omitempty"`
	NameKr        string `bson:"name_kr,omitempty" json:"name_kr,omitempty"`
	DescriptionEn string `bson:"description_en,omitempty" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr,omitempty" json:"description_kr,omitempty"`
}

type ListKPIParentResponse struct {
	Count int          `json:"count,omitempty"`
	Items []*KPIParent `json:"items,omitempty"`
}

type ListKPIParentRequest struct {
	Filter
	Year string `json:"year"`
}

func (k *KPIParent) BeforeCreate() {
	now := time.Now()
	k.CreatedAt = now
	k.UpdatedAt = now
}
