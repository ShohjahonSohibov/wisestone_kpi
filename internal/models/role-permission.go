package models

import (
	"time"
)

type RolePermission struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"`
	RoleId       string    `bson:"role_id" json:"role_id" binding:"required"`
	PermissionId string    `bson:"permission_id" json:"permission_id" binding:"required"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at"`
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

// BeforeCreate sets timestamps before creating a record
func (u *RolePermission) BeforeCreate() {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
}
