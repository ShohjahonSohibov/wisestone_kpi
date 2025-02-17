package models

import (
	"time"
)

type Permission struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	ActionKr      string    `bson:"action_kr" json:"action_kr" binding:"required"`
	ActionEn      string    `bson:"action_en" json:"action_en"`
	DescriptionKr string    `bson:"description_kr" json:"description_kr"`
	DescriptionEn string    `bson:"description_en" json:"description_en"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreatePermission struct {
	ActionKr      string `bson:"action_kr" json:"action_kr" binding:"required"`
	ActionEn      string `bson:"action_en" json:"action_en"`
	DescriptionKr string `bson:"description_kr" json:"description_kr"`
	DescriptionEn string `bson:"description_en" json:"description_en"`
}

type UpdatePermission struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	ActionKr      string `bson:"action_kr" json:"action_kr" binding:"required"`
	ActionEn      string `bson:"action_en" json:"action_en"`
	DescriptionKr string `bson:"description_kr" json:"description_kr"`
	DescriptionEn string `bson:"description_en" json:"description_en"`
}

type ListPermissionResponse struct {
	Count       int           `json:"count"`
	Items []*Permission `json:"items"`
}

type ListPermissionRequest struct {
	Filter
}

// BeforeCreate sets timestamps before creating a record
func (p *Permission) BeforeCreate() {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
}
