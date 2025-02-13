package models

import (
	"time"
)

type Permission struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	ActionKr      string    `bson:"action_kr" json:"action_kr" binding:"required"`
	ActionRu      string    `bson:"action_ru" json:"action_ru"`
	DescriptionKr string    `bson:"description_kr" json:"description_kr"`
	DescriptionRu string    `bson:"description_ru" json:"description_ru"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreatePermission struct {
	ActionKr      string `bson:"action_kr" json:"action_kr" binding:"required"`
	ActionRu      string `bson:"action_ru" json:"action_ru"`
	DescriptionKr string `bson:"description_kr" json:"description_kr"`
	DescriptionRu string `bson:"description_ru" json:"description_ru"`
}

type UpdatePermission struct {
	ID            string `json:"id,omitempty" bson:"_id,omitempty"`
	ActionKr      string `bson:"action_kr" json:"action_kr" binding:"required"`
	ActionRu      string `bson:"action_ru" json:"action_ru"`
	DescriptionKr string `bson:"description_kr" json:"description_kr"`
	DescriptionRu string `bson:"description_ru" json:"description_ru"`
}

type ListPermissionResponse struct {
	Count       int           `json:"count"`
	Permissions []*Permission `json:"permissions"`
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
