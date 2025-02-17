package models

import "time"

type KPIParent struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string    `bson:"name_en" json:"name_en"`
	NameKr        string    `bson:"name_kr" json:"name_kr"`
	DescriptionEn string    `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string    `bson:"description_kr" json:"description_kr,omitempty"`
	Year          string    `bson:"year" json:"year"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateKPIParent struct {
	NameEn        string `bson:"name_en" json:"name_en"`
	NameKr        string `bson:"name_kr" json:"name_kr"`
	DescriptionEn string `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr" json:"description_kr,omitempty"`
	Year          string `bson:"year" json:"year"`
}

type UpdateKPIParent struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string `bson:"name_en" json:"name_en"`
	NameKr        string `bson:"name_kr" json:"name_kr"`
	DescriptionEn string `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr" json:"description_kr,omitempty"`
	Year          string `bson:"year" json:"year"`
}

type ListKPIParentResponse struct {
	Count int          `json:"count"`
	Items []*KPIParent `json:"items"`
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
