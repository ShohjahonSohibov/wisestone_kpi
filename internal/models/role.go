package models

import (
	"time"
)

type Role struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn        string    `bson:"name_en" json:"name_en" binding:"required"`
	NameKr        string    `bson:"name_kr" json:"name_kr" binding:"required"`
	CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateRole struct {
	NameEn      string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr      string `bson:"name_kr" json:"name_kr" binding:"required"`
}

type UpdateRole struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn      string `bson:"name_en" json:"name_en" binding:"required"`
	NameKr      string `bson:"name_kr" json:"name_kr" binding:"required"`
}

type ListRoleResponse struct {
	Count int     `json:"count"`
	Items []*Role `json:"items"`
}

type ListRoleRequest struct {
	Filter
}

// BeforeCreate sets timestamps before creating a record
func (u *Role) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}
