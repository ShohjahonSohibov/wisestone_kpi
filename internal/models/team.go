package models

import (
	"time"
)

type Team struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	NameUz        string    `bson:"name_uz" json:"name_uz" binding:"required"`
	NameEn        string    `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string    `bson:"name_kr" json:"name_kr" binding:"required"`
	DescriptionUz string    `bson:"description_uz" json:"description_uz,omitempty"`
	DescriptionEn string    `bson:"description_en" json:"description_en,omitempty"`
	DescriptionKr string    `bson:"description_kr" json:"description_kr,omitempty"`
	LeaderId      string    `bson:"leader_id" json:"leader_id" binding:"required"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateTeam struct {
	NameUz      string `bson:"name_uz" json:"name_uz" binding:"required"`
	NameEn      string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr      string `bson:"name_kr" json:"name_kr" binding:"required"`
	Description string `json:"description,omitempty"`
	LeaderId    string `json:"leader_id" binding:"required"`
}

type UpdateTeam struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	NameUz      string `bson:"name_uz" json:"name_uz" binding:"required"`
	NameEn      string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr      string `bson:"name_kr" json:"name_kr" binding:"required"`
	Description string `bson:"description" json:"description,omitempty"`
	LeaderId    string `bson:"leader_id" json:"leader_id" binding:"required"`
}

type ListTeamsResponse struct {
	Count int     `json:"count"`
	Teams []*Team `json:"teams"`
}

type ListTeamsRequest struct {
	Filter Filter `json:"filter"`
}

// BeforeCreate sets timestamps before creating a record
func (u *Team) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}
