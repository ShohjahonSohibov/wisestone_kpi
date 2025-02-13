package models

import (
	"time"
)

type RolePermission struct {
	ID           string          `json:"id,omitempty" bson:"_id,omitempty"`
	RoleId       string          `bson:"role_id,omitempty" json:"role_id,omitempty"`
	Role         ShortRole       `bson:"role" json:"role"`
	PermissionId string          `bson:"permission_id,omitempty" json:"permission_id,omitempty"`
	Permission   ShortPermission `bson:"permission" json:"permission"`
	CreatedAt    time.Time       `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateRolePermission struct {
	RoleId       string `bson:"role_id" json:"role_id" binding:"required"`
	PermissionId string `bson:"permission_id" json:"permission_id" binding:"required"`
}

type UpdateRolePermission struct {
	ID           string `json:"id,omitempty" bson:"_id,omitempty"`
	RoleId       string `bson:"role_id" json:"role_id" binding:"required"`
	PermissionId string `bson:"permission_id" json:"permission_id" binding:"required"`
}

type ListRolePermissionResponse struct {
	Count int               `json:"count"`
	Data  []*RolePermission `json:"data"`
}

type ListRolePermissionRequest struct {
	RoleId       string `bson:"role_id" json:"role_id"`
	PermissionId string `bson:"permission_id" json:"permission_id"`
	Filter
}

type ShortRole struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	NameEn string `bson:"name_en,omitempty" json:"name_en,omitempty" `
	NameKr string `bson:"name_kr,omitempty" json:"name_kr,omitempty" `
}

type ShortPermission struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	ActionKr string `bson:"action_kr,omitempty" json:"action_kr,omitempty"`
	ActionRu string `bson:"action_ru,omitempty" json:"action_ru,omitempty"`
}

// BeforeCreate sets timestamps before creating a record
func (u *RolePermission) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}
