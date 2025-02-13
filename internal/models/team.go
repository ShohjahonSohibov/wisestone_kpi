package models

import (
	"time"
)

type Team struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string    `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string    `bson:"name_kr" json:"name_kr" binding:"required"`
	DescriptionEn string    `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string    `bson:"description_kr" json:"description_kr,omitempty"`
	LeaderId      string    `bson:"leader_id" json:"leader_id"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateTeam struct {
	NameEn        string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string `bson:"name_kr" json:"name_kr" binding:"required"`
	DescriptionEn string `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr" json:"description_kr,omitempty"`
	LeaderId      string `bson:"leader_id" json:"leader_id"`
}

type UpdateTeam struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string `bson:"name_kr" json:"name_kr" binding:"required"`
	DescriptionEn string `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string `bson:"description_kr" json:"description_kr,omitempty"`
	LeaderId      string `bson:"leader_id" json:"leader_id"`
}

type ListTeamsResponse struct {
	Count int     `json:"count"`
	Items []*Team `json:"items"`
}

type ListTeamsRequest struct {
	Filter
}

// BeforeCreate sets timestamps before creating a record
func (u *Team) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}
